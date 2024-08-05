const std = @import("std");
const zap = @import("zap");

//JSON;
const Value = std.json.Value;
const parseFromSlice = std.json.parseFromSlice;

// const STORAGE = "/var/lib/files";
const STORAGE = "/tmp";
// const CODE_STORAGE = STORAGE ++ "/code";
// const OUT_STORAGE = STORAGE ++ "/output";
const FILE_NAME_LEN = 100;
const RESPONSE_LEN = 2000;

//Timestamp
const time = @import("datetime.zig");

const Response = struct {
    output: ?[]const u8 = null,
    err: ?[]const u8 = null,
    timestamp: ?[]const u8 = null,
};

fn on_request(r: zap.Request) void {
    // parse for FORM (body) parameters first
    r.parseBody() catch |err| {
        std.log.err("Parse Body error: {any}. Expected if body is empty", .{err});
    };
    if (r.body) |body| {
        const allocator = std.heap.page_allocator;

        std.log.info("Body length is {any}\n", .{body.len});
        std.debug.print("Body: {s}\n", .{body});

        // Parse JSON
        // https://github.com/ziglang/zig/blob/master/lib/std/json/dynamic_test.zig
        var parsed = parseFromSlice(Value, allocator, body, .{}) catch return;
        defer parsed.deinit();
        const root = parsed.value;
        const data = root.object.get("data").?;
        std.debug.print("Data: '{s}' {}\n", .{ data.string, data.string.len });

        const id: i32 = 1234;
        const extension = "py";
        var code_buf = [_]u8{undefined} ** FILE_NAME_LEN;
        const code_id = std.fmt.bufPrint(&code_buf, "{s}/code_{d}.{s}", .{ STORAGE, id, extension }) catch return;
        var out_buf = [_]u8{undefined} ** FILE_NAME_LEN;
        const out_id = std.fmt.bufPrint(&out_buf, "{s}/output_{d}", .{ STORAGE, id }) catch return;
        var err_buffer = [_]u8{undefined} ** FILE_NAME_LEN;
        const err_id = std.fmt.bufPrint(&err_buffer, "{s}/error_{d}", .{ STORAGE, id }) catch return;

        // var dir = std.fs.openDirAbsolute(STORAGE, .{}) catch return;
        // defer dir.close();
        var code_file = std.fs.createFileAbsolute(code_id, .{}) catch return;
        defer code_file.close();

        // Write data to file
        const comment = "# ";
        var code_stream = std.io.bufferedWriter(code_file.writer());
        const code_writer = code_stream.writer();
        code_writer.print("{s} begin\n", .{comment}) catch return;
        code_writer.writeAll(data.string) catch return;
        code_writer.print("\n{s} end", .{comment}) catch return;
        code_stream.flush() catch return;

        // zap/zig-linux-x86_64-0.13.0/lib/std/posix.zig
        // test "pipe" {
        //     if (native_os == .windows or native_os == .wasi)
        //         return error.SkipZigTest;

        //     const fds = try posix.pipe();
        //     try expect((try posix.write(fds[1], "hello")) == 5);
        //     var buf: [16]u8 = undefined;
        //     try expect((try posix.read(fds[0], buf[0..])) == 5);
        //     try testing.expectEqualSlices(u8, buf[0..5], "hello");
        //     posix.close(fds[1]);
        //     posix.close(fds[0]);
        // }

        // zap/zig-linux-x86_64-0.13.0/lib/std/os/linux.zig
        const pid: i32 = @intCast(std.os.linux.fork());

        if (pid == 0) {
            std.debug.print("Child process: {}\n", .{pid});

            // STDOUT
            var out_fd = std.fs.createFileAbsolute(out_id, .{}) catch return;
            defer out_fd.close();
            std.posix.dup2(out_fd.handle, std.posix.STDOUT_FILENO) catch |err| {
                std.debug.print("Error: {}\n", .{err});
                std.os.linux.exit(1);
            };

            // STDERR
            var fd_err = std.fs.createFileAbsolute(err_id, .{}) catch return;
            defer fd_err.close();
            std.posix.dup2(fd_err.handle, std.posix.STDERR_FILENO) catch |err| {
                std.debug.print("Error: {}\n", .{err});
                std.os.linux.exit(1);
            };

            const argv = [_][]const u8{ "python3", code_id };

            // /lib/std/process.zig
            const error_code = std.process.execv(allocator, &argv);

            std.debug.print("ERROR: {}\n", .{error_code});

            // pub const ExecveError = error{
            //     SystemResources,
            //     AccessDenied,
            //     InvalidExe,
            //     FileSystem,
            //     IsDir,
            //     FileNotFound,
            //     NotDir,
            //     FileBusy,
            //     ProcessFdQuotaExceeded,
            //     SystemFdQuotaExceeded,
            //     NameTooLong,
            // } || UnexpectedError;

            std.os.linux.exit(1);
        } else {
            var status: u32 = 0;

            std.debug.print("children pid: {}\n", .{pid});

            _ = std.os.linux.waitpid(pid, &status, 0);
            if (status != 0) {
                std.debug.print("Process execution error {d}", .{status});
                // var err_msg_buf = [_]u8{undefined} ** 100;
                // const err_msg = std.fmt.bufPrint(&err_msg_buf, "Process execution error {d}", .{status}) catch return;
                // r.setContentType(.TEXT) catch return;
                // r.sendBody(err_msg) catch return;
                // return;
            }

            // Write output file content to response
            var out_fd = std.fs.openFileAbsolute(out_id, .{}) catch return;
            defer out_fd.close();
            const out_size = out_fd.getEndPos() catch return;
            var out_stream = std.io.bufferedReader(out_fd.reader());
            const out_reader = out_stream.reader();
            const out_txt = out_reader.readAllAlloc(allocator, out_size) catch return;
            defer allocator.free(out_txt);

            // Write errors file content to response
            var err_fd = std.fs.openFileAbsolute(err_id, .{}) catch return;
            defer err_fd.close();
            const err_size = err_fd.getEndPos() catch return;
            var err_stream = std.io.bufferedReader(err_fd.reader());
            const err_reader = err_stream.reader();
            const err_txt = err_reader.readAllAlloc(allocator, err_size) catch return;
            defer allocator.free(err_txt);

            // Timestamp
            const now = time.Datetime.now();
            const now_str = now.formatHttp(allocator) catch return;
            defer allocator.free(now_str);

            const res: Response = .{ .output = out_txt, .err = err_txt, .timestamp = now_str };

            var buf: [RESPONSE_LEN]u8 = undefined;
            if (zap.stringifyBuf(&buf, res, .{})) |json| {
                r.setContentType(.JSON) catch return;
                r.sendBody(json) catch return;
            } else {
                r.setContentType(.TEXT) catch return;
                r.sendBody("Json parser error: Output too long") catch return;
            }
        }
    }
    // parse potential query params (for ?terminate=true)
    // r.parseQuery();
    // const param_count = r.getParamCount();
    // std.log.info("param_count: {}", .{param_count});

    // const alloc: std.mem.Allocator = undefined;

    // const params = r.parametersToOwnedList(alloc, false) catch unreachable;
    // defer params.deinit();
    // for (params.items) |kv| {
    //     if (kv.value) |v| {
    //         std.debug.print("\n", .{});
    //         std.log.info("Param `{s}` in owned list is {any}\n", .{ kv.key.str, v });
    //         if (r.getParamStr(alloc, kv.key.str, false)) |maybe_str| {
    //             const value: []const u8 = if (maybe_str) |s| s.str else "(no value)";
    //             // above, we didn't defer s.deinit because the string is just a slice from the request buffer
    //             std.log.debug("   {s} = {s}", .{ kv.key.str, value });
    //         } else |err| {
    //             std.log.err("Error: {any}\n", .{err});
    //         }
    //     }
    // }

    return;
}

pub fn main() !void {
    var listener = zap.HttpListener.init(.{
        .port = 4000,
        .on_request = on_request,
        .log = false,
    });
    try listener.listen();

    std.debug.print(
        \\ Listening on 0.0.0.0:4000
        \\ 
        \\ Check out:
        \\ http://localhost:4000/user/1   # -- first user
        \\ http://localhost:4000/user/2   # -- second user
        \\ http://localhost:4000/user/3   # -- non-existing user
        \\
    , .{});

    // Start worker threads
    zap.start(.{
        .threads = 1,
        .workers = 1, // User map cannot be shared among multiple worker processes
    });
}

// test "simple test" {
//     var list = std.ArrayList(i32).init(std.testing.allocator);
//     defer list.deinit(); // try commenting this out and see if zig detects the memory leak!
//     try list.append(42);
//     try std.testing.expectEqual(@as(i32, 42), list.pop());
// }

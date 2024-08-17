import { renderToReadableStream } from "react-dom/server";

function Component(props: { message: string }) {
  return (
    <body>
      <h1>{props.message}</h1>
    </body>
  );
}

const server=Bun.serve({
  port: 3000,
  async fetch() {
    const stream = await renderToReadableStream(
      <Component message="Hello from server!" />,
    );
    return new Response(stream, {
      headers: { "Content-Type": "text/html" },
    });
  },
});

console.log(`Listening on http://localhost:${server.port} ...`);
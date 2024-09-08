import {
  // Form,
  // NavLink,
  Links,
  Meta,
  Form,
  // Outlet,
  // Scripts,
  // ScrollRestoration,
  useLoaderData,
  // useNavigation,
  // useNavigate,
} from "@remix-run/react";
// import Box from "../../components/box";
import Text from "../../components/text";
import {getSheet} from '../sheet.server';

import type { LoaderFunctionArgs } from "@remix-run/node";
import { json } from "@remix-run/node";
import invariant from "tiny-invariant";

export const loader = async ({
    params,
  }: LoaderFunctionArgs) => {
    invariant(params.slug, "Missing sheetId param");
    const match = params.slug.match(/^(\d+)/);
    if (!match) {
        // Handle error, maybe redirect to a 404 page
        return;
    }
    const id = match[1];
    const sheet = await getSheet(id);
    if (!sheet) {
        throw new Response("Not Found", { status: 404 });
      }
    console.log(sheet);
    return json({ sheet });
  };

export default function Sheet() {
//   const contact = {
//     first: "Your",
//     last: "Name",
//     avatar: "https://placekitten.com/200/200",
//     twitter: "your_handle",
//     notes: "Some notes",
//     favorite: true,
//   };
  const { sheet } = useLoaderData<typeof loader>();

  return (
    <div id="sheet">
      <h1>{sheet.title}</h1>
      {/* <Text table="sheet" id={sheet.id} field="title" value={sheet.title} format="text-4xl"/> */}
      {/* <Box /> */}
    </div>
  );
}

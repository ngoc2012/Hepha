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
import Box from "../../components/box";
import Text from "../../components/text";

// type SheetProps = {
//   id: number
//   title: string
//   description?: string
//   owner?: string
//   protection?: number
//   searchable?: boolean
//   exeTimes?: number
//   comments?: number
//   evaluations?: number
//   star?: number
//   createdDate?: string
//   modDate?: string
// }

// export default async function Sheet({ params }: { params: { slug: string } }) {

//   const match = params.slug.match(/^(\d+)/);
//   if (!match) {
//     // Handle error, maybe redirect to a 404 page
//     return;
//   }

//   const id = match[1];

//   // Fetch data from external API
//   const res = await fetch(`http://${process.env.GIN_ADDR}/sheet/${id}`, {
//     method: 'GET',
//     cache: 'no-cache'
//   });

//   const repo: SheetProps = await res.json();

//   // console.log(repo)
  
//   return <>
//     <Text table="sheet" id={id} field="title" value={repo.title} format="text-4xl"/>
//     <Box />
//   </>;
// }

import type { LoaderFunctionArgs } from "@remix-run/node";
import { json } from "@remix-run/node";
import { getContact } from "../data";
import invariant from "tiny-invariant";

export const loader = async ({
    params,
  }: LoaderFunctionArgs) => {
    invariant(params.slug, "Missing contactId param");
    const contact = await getContact(params.slug);
    if (!contact) {
        throw new Response("Not Found", { status: 404 });
      }
    return json({ contact });
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
  const { contact } = useLoaderData<typeof loader>();

  return (
    <html lang="en">
      <head>
        <meta charSet="utf-8" />
        <meta name="viewport" content="width=device-width, initial-scale=1" />
        <Meta />
        <Links />
      </head>
      <body>
      </body>
      </html>
  );
}
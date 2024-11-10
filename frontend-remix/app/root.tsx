import type { LinksFunction } from "@remix-run/node";
import { json, redirect } from "@remix-run/node";
import {newSheet} from './sheet.server';
import {
  // Form,
  // Link,
  Links,
  Meta,
  Outlet,
  Scripts,
  // ScrollRestoration,
  useLoaderData,
  // useNavigation,
  // useNavigate,
} from "@remix-run/react";
import stylesheet from "~/tailwind.css?url";
// import { createEmptyContact, getContacts } from "./data";

function generateSlug(title: string): string {
  // Convert to lowercase
  let slug = title.toLowerCase();

  // Replace spaces with hyphens
  slug = slug.replace(/ /g, '-');

  // Remove special characters
  slug = slug.replace(/[^a-z0-9-]/g, '');

  // Remove multiple hyphens
  slug = slug.replace(/-+/g, '-');

  // Trim leading and trailing hyphens
  slug = slug.replace(/^-+|-+$/g, '');

  return slug;
}

export const action = async () => {
  // const contact = await createEmptyContact();
  // return redirect(`/contacts/${contact.id}/edit`);
};

export const links: LinksFunction = () => [
  { rel: "stylesheet", href: stylesheet },
];

// export const loader = async () => {
  // const contacts = await getContacts();
  // const sheets = await getSheets(10, 1)
  // return json({ sheets });
// };

export default function App() {
  // const navigate = useNavigate();

  return (
    <html lang="en">
      <head>
        <meta charSet="utf-8" />
        <meta name="viewport" content="width=device-width, initial-scale=1" />
        <Meta />
        <Links />
      </head>
      <body>
        <nav className="flex">
          <h1>Remix Sheets</h1>
          <div>
            <form method="post">
              <button
                type="submit"
                className="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded"
              >
                New Sheet
              </button>
            </form>
          </div>
        </nav>
        <div id="main">
          <Outlet />
        </div>
        <Scripts />
      </body>
    </html>
  );
}

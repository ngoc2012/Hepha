import type { LinksFunction } from "@remix-run/node";
import { json, redirect } from "@remix-run/node";
import {getSheets, newSheet} from '../sheet.server';
import {
  // Form,
  // Link,
  Links,
  Meta,
  // Outlet,
  // Scripts,
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

export const loader = async () => {
  // const contacts = await getContacts();
  const sheets = await getSheets(10, 1)
  return json({ sheets });
};

export default function App() {
  const { sheets } = useLoaderData<typeof loader>();
  // const navigate = useNavigate();

  return (
    <div id="sheets list">
      {sheets.length ? (
      <table className="min-w-full divide-y divide-gray-200">
        <thead className="bg-gray-50">
          <tr>
            <th scope="col" className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
              Name
            </th>
            <th scope="col" className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
              Id
            </th>
            <th scope="col" className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
              Description
            </th>
            <th scope="col" className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
              Author
            </th>
          </tr>
        </thead>
        <tbody className="bg-white divide-y divide-gray-200">
          {sheets.map((sheet: any) => (
          <tr className="hover:bg-slate-100"
            key={sheet.id}
            // onClick={handleSheetClick(sheet.id, sheet.title)}
            >
            
            <td className="px-6 py-4 whitespace-nowrap">
              <div className="flex items-center">
                <div className="ml-4">
                  <div className="text-sm font-medium hover:font-extrabold text-gray-900">
                  <a href={`/sheet/${sheet.id}-${generateSlug(sheet.title)}`}>
                  {sheet.title}
                  </a>
                  </div>
                </div>
              </div>
            </td>
            <td className="px-6 py-4 whitespace-nowrap">
              <div className="text-sm text-gray-900">{sheet.id}</div>
            </td>
            <td className="px-6 py-4 whitespace-nowrap">
              <div className="text-sm text-gray-900">{sheet.description}</div>
            </td>
            <td className="px-6 py-4 whitespace-nowrap">
              <div className="text-sm text-gray-900">{sheet.owner}</div>
            </td>
            
          </tr>
          ))}
        </tbody>
      </table>

      ) : (
        <p>
          <i>No sheets</i>
        </p>
      )}
    </div>
  );
}

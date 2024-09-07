"use client" 

import { useEffect, useState } from 'react';
import { useRouter } from 'next/navigation'

import {getSheets, newSheet} from '../SSR/sheet';

interface Sheet {
  id: number;
  title: string;
  description: string;
  owner: string;
  protection: number;
  searchable: boolean;
  exe_times: number;
  comments: number;
  evaluations: number;
  star: number;
  created_date: Date;
  mod_date: Date;
}

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


export default function Home() {
  const router = useRouter()
  const [sheets, setSheets] = useState<Sheet[]>([])

  useEffect(() => {
    async function fetchSheets() {
      setSheets(await getSheets(10, 1));
      // console.log(sheets)
    }
    fetchSheets();
  }, []);

  const handleNewSheetButtonClick = async () => {
    const res = await newSheet("hepha", "No title", "No description")
    if (res.error) {
      console.error(res.error)
      return;
    }
    // console.log(res)
    router.push('/sheet/' + res.slug)
  };

  const handleSheetClick = (id: number, title: string) => () => {
    router.push('/sheet/' + id + '-' + generateSlug(title))
  };

  // https://tailwindcss.com/docs/hover-focus-and-other-states
  return (
    <div>
        {/* <Navbar /> */}
        {/* <h2 className="text-2xl font-semibold">Home</h2> */}
        <button
          className="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded"
          // onClick={newButton}
          onClick={handleNewSheetButtonClick}
        >
          New Sheet
        </button>
        <h1>Sheet List</h1>
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
            {sheets.map((sheet) => (
            <tr className="hover:bg-slate-100" key={sheet.id} onClick={handleSheetClick(sheet.id, sheet.title)}>
              <td className="px-6 py-4 whitespace-nowrap">
                <div className="flex items-center">
                  <div className="ml-4">
                    <div className="text-sm font-medium text-gray-900">
                    {sheet.title}
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
    </div>
  );
}

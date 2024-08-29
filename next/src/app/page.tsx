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
        <ul>
          {sheets.map((sheet) => (
            <li key={sheet.id}>
              <div className="flex space-x-4" onClick={handleSheetClick(sheet.id, sheet.title)}>
                <h2>{sheet.title}</h2>
                <p>{sheet.id}</p>
                <p>{sheet.description}</p>
                <p> by {sheet.owner}</p>
                {/* <p>Protection: {sheet.protection}</p>
                <p>Searchable: {sheet.searchable ? 'Yes' : 'No'}</p>
                <p>Executed Times: {sheet.exe_times}</p>
                <p>Comments: {sheet.comments}</p>
                <p>Evaluations: {sheet.evaluations}</p>
                <p>Star Rating: {sheet.star}</p>
                <p>Created Date: {new Date(sheet.created_date).toLocaleString()}</p>
                <p>Modified Date: {new Date(sheet.mod_date).toLocaleString()}</p> */}
              </div>
            </li>
          ))}
        </ul>
    </div>
  );
}

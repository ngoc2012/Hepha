"use client" 

import { useEffect, useState } from 'react';
import { Sheet } from '../types/sheet';
import { redirect } from 'next/navigation'

import Box from  "../components/box";

import getSheets from '../getSheets';
import newSheet from '@/newSheet';

interface Props {
  sheets: Sheet[];
}

export default function Home() {
  const [sheets, setSheets] = useState<Sheet[]>([]);

  useEffect(() => {
    async function fetchSheets() {
      setSheets(await getSheets(10, 1));
      console.log(sheets);
    }
    fetchSheets();
  }, []);

  const handleNewSheetButtonClick = async () => {
    const res = await newSheet("hepha");
    if (res.error) {
      console.error(res.error);
      return;
    }
    redirect('/sheet/' + res.id + res.slug);
  };

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
              <div className="flex space-x-4">
                <h2>{sheet.title}</h2>
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
        <Box />
    </div>
  );
}

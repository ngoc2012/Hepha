import { GetServerSideProps } from 'next';
import { Sheet } from '../types/sheet';

// import Navbar from "../components/navbar";
import Box from  "../components/box";

interface Props {
  sheets: Sheet[];
}

export default function Home({ sheets }: Props) {
  return (
    <div>
        {/* <Navbar /> */}
        {/* <h2 className="text-2xl font-semibold">Home</h2> */}
        <button
          className="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded"
          // onClick={newButton}
        >
          New
        </button>
        <h1>Sheet List</h1>
        <ul>
          {sheets.map((sheet) => (
            <li key={sheet.id}>
              <h2>{sheet.title}</h2>
              <p>{sheet.description}</p>
              <p>Owner: {sheet.owner}</p>
              <p>Protection: {sheet.protection}</p>
              <p>Searchable: {sheet.searchable ? 'Yes' : 'No'}</p>
              <p>Executed Times: {sheet.exe_times}</p>
              <p>Comments: {sheet.comments}</p>
              <p>Evaluations: {sheet.evaluations}</p>
              <p>Star Rating: {sheet.star}</p>
              <p>Created Date: {new Date(sheet.created_date).toLocaleString()}</p>
              <p>Modified Date: {new Date(sheet.mod_date).toLocaleString()}</p>
            </li>
          ))}
        </ul>
        <Box />
    </div>
  );
}

export const getServerSideProps: GetServerSideProps = async () => {
  // Replace with your actual API endpoint
  // const res = await fetch('https://api.example.com/sheets');
  

  const res = await fetch('http://' + process.env.BACKEND_ADDR + '/sheets', {
    method: 'POST',
    headers: {
      'Content-Type': 'text/plain',
    },
    body: JSON.stringify({ name: "", length: 10, page : 1 }),
    cache: 'no-cache'
  });

  const sheets: Sheet[] = await res.json();

  return {
    props: {
      sheets,
    },
  };
};

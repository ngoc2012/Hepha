'use server'

export default async function getSheets(length: number, page: number) {

  const res = await fetch('http://' + process.env.GIN_ADDR + '/sheets', {
    method: 'POST',
    headers: {
      'Content-Type': 'text/plain',
    },
    body: JSON.stringify({ name: "", length: length, page : page }),
    cache: 'no-cache'
  });
 
  if (!res.ok) {
    throw new Error('Failed to fetch data')
  }

  const json = await res.json();
  // console.log(json);
  if (json.error) {
    console.error(json.error);
    return [];
  }

  return json
}
'use server'

export default async function getSheet(id: number) {

  const res = await fetch('http://' + process.env.GIN_ADDR + '/sheet/' + id.toString(), {
    method: 'GET',
    headers: {
      'Content-Type': 'text/plain',
    },
    cache: 'no-cache'
  });

  // console.log("getSheets", res)
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
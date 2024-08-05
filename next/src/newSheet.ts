'use server'

export default async function newSheet(user: string) {

  // const res = await fetch('http://'+ url + ':4000/user?id=1', {
  const res = await fetch('http://' + process.env.BACKEND_ADDR + '/new_sheet', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify({ user: user }),
    cache: 'no-cache'
  });
 
  if (!res.ok) {
    // This will activate the closest `error.js` Error Boundary
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
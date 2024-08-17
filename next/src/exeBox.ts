'use server'

export default async function exeBox(data_type:string, data: string, lang: string, ver: string) {
//   const res = await fetch('https://api.example.com/...')
  // The return value is *not* serialized
  // You can return Date, Map, Set, etc.

  // const res = await fetch('http://'+ url + ':4000/user?id=1', {
  const res = await fetch('http://' + process.env.GIN_ADDR + '/new_box', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify({ user: "hepha", sheet: 1, type: data_type, lang: lang, ver: ver, data: data, caption: "test" }),
    cache: 'no-cache'
  });
 
  if (!res.ok) {
    // This will activate the closest `error.js` Error Boundary
    throw new Error('Failed to fetch data')
  }
 
  return res.text()
}
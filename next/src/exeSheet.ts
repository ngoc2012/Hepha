'use server'

export default async function exeSheet(data: string) {

  // const res = await fetch('http://'+ url + ':4000/user?id=1', {
  const res = await fetch('http://' + process.env.BACKEND_ADDR + '/new_box', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify({ user: "hepha", sheet: 1, type: "Code", language: "Python", language_version: "3", data: data, caption: "test" }),
    cache: 'no-cache'
  });
 
  if (!res.ok) {
    // This will activate the closest `error.js` Error Boundary
    throw new Error('Failed to fetch data')
  }
 
  return res.text()
}
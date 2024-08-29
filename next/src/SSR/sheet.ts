'use server'

export async function getSheets(length: number, page: number) {

  const res = await fetch('http://' + process.env.GIN_ADDR + '/sheets', {
    method: 'POST',
    headers: {
      'Content-Type': 'text/plain',
    },
    body: JSON.stringify({ name: "", length: length, page : page }),
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

export async function exeSheet(data: string) {

  // const res = await fetch('http://'+ url + ':4000/user?id=1', {
  const res = await fetch('http://' + process.env.GIN_ADDR + '/new_box', {
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

export async function newSheet(owner: string, title: string, description: string) {

  // const res = await fetch('http://'+ url + ':4000/user?id=1', {
  const res = await fetch('http://' + process.env.GIN_ADDR + '/new_sheet', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify({ owner: owner, title: title, description: description }),
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
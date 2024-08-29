'use server'

export async function editText(table: string, id: string, field: string, value: string) {

  const res = await fetch('http://' + process.env.GIN_ADDR + '/edit/' + table + '/' + id, {
    method: 'POST',
    body: JSON.stringify({ [field]: value }),
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
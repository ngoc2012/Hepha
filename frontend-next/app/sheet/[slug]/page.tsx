import Box from "@/components/box";
import Text from "@/components/text";

type SheetProps = {
  id: number
  title: string
  description?: string
  owner?: string
  protection?: number
  searchable?: boolean
  exeTimes?: number
  comments?: number
  evaluations?: number
  star?: number
  createdDate?: string
  modDate?: string
}

export default async function Sheet({ params }: { params: { slug: string } }) {

  const match = params.slug.match(/^(\d+)/);
  if (!match) {
    // Handle error, maybe redirect to a 404 page
    return;
  }

  const id = match[1];

  // Fetch data from external API
  const res = await fetch(`http://${process.env.GIN_ADDR}/sheet/${id}`, {
    method: 'GET',
    cache: 'no-cache'
  });

  const repo: SheetProps = await res.json();

  // console.log(repo)
  
  return <>
    <Text table="sheet" id={id} field="title" value={repo.title} format="text-4xl"/>
    <Box />
  </>;
}
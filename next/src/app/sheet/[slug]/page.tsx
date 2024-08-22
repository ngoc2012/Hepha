import type { InferGetServerSidePropsType, GetServerSideProps } from 'next'
import Box from "@/components/box";

type SheetProps = {
  id: number
  slug: string
}

export const getServerSideProps = (async (context) => {
  if (!context.params || !context.params.slug  || typeof context.params.slug !== 'string') {
    return { notFound: true }
  }

  const match = context.params.slug.match(/^(\d+)/)
  if (!match) {
    return { notFound: true }
  }

  const id = match[1]

  // Fetch data from external API
  const res = await fetch(`http://${process.env.GIN_ADDR}/sheet/${id}`, {
    method: 'GET',
    cache: 'no-cache'
  });
  const repo: SheetProps = await res.json()
  return { props: { repo } }
}) satisfies GetServerSideProps<{ repo: SheetProps }>

export default function Sheet({ repo,}: InferGetServerSidePropsType<typeof getServerSideProps>) {
  return <>
    <div>{repo.id}</div>
    <div>My Sheet : {repo.slug}</div>
    <Box />
  </>
}
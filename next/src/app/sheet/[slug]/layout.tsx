import localFont from "next/font/local";

const geistSans = localFont({
  src: "./fonts/GeistVF.woff",
  variable: "--font-geist-sans",
});
const geistMono = localFont({
  src: "./fonts/GeistMonoVF.woff",
  variable: "--font-geist-mono",
});

export default function SheetLayout({
  children,
}: {
  children: React.ReactNode
}) {
  return <section  className={`${geistSans.variable} ${geistMono.variable}`}>{children}</section>
}
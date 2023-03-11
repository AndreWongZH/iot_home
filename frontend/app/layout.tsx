import './globals.css'

export default function RootLayout({
  children,
}: {
  children: React.ReactNode
}) {
  return (
    <html lang="en">
      {/*
        <head /> will contain the components returned by the nearest parent
        head.tsx. Find out more at https://beta.nextjs.org/docs/api-reference/file-conventions/head
      */}
      <head />
      <body className='bg-outbounds'>
        <div className='max-w-md mx-auto bg-background-default min-h-screen h-max'>
          {children}
        </div>
      </body>
    </html>
  )
}

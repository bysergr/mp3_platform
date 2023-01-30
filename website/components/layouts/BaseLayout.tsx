import Head from 'next/head'
import Footer from "../widgets/Footer"
import Header from "../widgets/Header"

const BaseLayout = ({ children } : { children: any }) => {
  return (
    <>
      <Head>
        <title>MP3 PLATFORM</title>
        <meta name="description" content="WebSite for the project mp3 platform" />
        <meta name="viewport" content="width=device-width, initial-scale=1" />
        <link rel="icon" href="/favicon.ico" />
      </Head>
      <Header />
        <main className="px-8 min-h-[calc(100vh-8.5rem)] flex justify-center align-middle">
          {children}
        </main>
      <Footer />
    </>
  )
}

export default BaseLayout

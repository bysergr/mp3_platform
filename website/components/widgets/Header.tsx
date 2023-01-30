import Link from "next/link"

const Header = () => {
  return (
    <header className="flex justify-between h-20 max-w-2xl px-8 mx-auto">
      <Link href="/" className="my-auto font-semibold text-lg">
        Home
      </Link> 
      <div className="my-auto flex gap-4">
        <Link href="/" className="my-auto font-semibold text-lg">
          Source 
        </Link> 
      </div>
    </header>
  )
} 

export default Header

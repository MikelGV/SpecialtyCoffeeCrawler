import Link from "next/link";

export default function TopBar() {
    return (
        <nav className="flex font-pacifico bg-green-950 px-4 border-b md:shadow-lg items-center relative">
            <div className="text-lg md:py-0 py-4">
                Specialty Coffee Europe
            </div>
            <ul className="md:px-2 ml-auto md:flex md:space-x-2 absolute md:relative top-full left-0 right-0">
                <li>
                    <Link href="/" className="flex md:inline-flex p-4 items-center hover:bg-green-955">
                        <span>Home</span>
                    </Link>
                </li>
                <li>
                    <Link href="/products" className="flex md:inline-flex p-4 items-center hover:bg-green-955">
                        <span>Products</span>
                    </Link>
                </li>
            </ul>
        </nav>
            
    )
}

import Link from "next/link";

export default function TopBar() {
    return (
        <nav className="flex font-pacifico bg-green-950 px-4 border-b md:shadow-lg items-center justify-between">
            <h1 className="text-lg py-4">
                Specialty Coffee Europe
            </h1>
            <ul className="flex items-center space-x-6">
                <li>
                    <Link href="/" className="p-4 hover:bg-green-955">
                        <span>Home</span>
                    </Link>
                </li>
                <li>
                    <Link href="/products" className="p-4 hover:bg-green-955">
                        <span>Products</span>
                    </Link>
                </li>
            </ul>
        </nav>
            
    )
}

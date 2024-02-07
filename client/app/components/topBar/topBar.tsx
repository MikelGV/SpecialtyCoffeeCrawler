import Link from "next/link";

export default function TopBar() {
    return (
        <div className="flex-col">
            <div className="bg-yellow-950 p-0.5"></div>
            <div className="flex bg-green-950 p-4 rounded-b-xl font-pacifico">
                <div id="left" className="flex-1 pl-8">
                    <Link className="" href="/">Home</Link>
                </div>
                <div id="right" className="pr-8">
                    <Link className="" href="/products">Products</Link>
                </div>
            </div>
        </div>
    )
}

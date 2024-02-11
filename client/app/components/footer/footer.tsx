import Link from "next/link";

export default function Footer() {
    return (
        <div className="flex-col">
            <div className="flex bg-green-950 p-10 font-pacifico justify-center border-t md:shadow-lg">
                <Link className="" href="https://mikelgv.com" target="_blank">Made by MikelGV</Link>
            </div>
        </div>
    )
}

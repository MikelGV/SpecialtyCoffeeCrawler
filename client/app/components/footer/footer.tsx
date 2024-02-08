import Link from "next/link";

export default function Footer() {
    return (
        <div className="flex-col">
            <div className="flex bg-green-950 p-10 rounded-t-xl font-pacifico justify-center">
                <Link className="" href="https://mikelgv.com" target="_blank">Made by MikelGV</Link>
            </div>
            <div className="bg-yellow-950 p-0.5"></div>
        </div>
    )
}

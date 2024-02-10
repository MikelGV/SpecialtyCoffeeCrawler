import Link from "next/link"
import { PlayIcon } from "@heroicons/react/16/solid"

export default function HomeS() {
    return (

        <div>
            <div className="flex justify-center m-56 mb-44">
                <h1 className="text-black font-pacifico text-6xl">Specialty Coffee Europe</h1>
            </div>
            <div className="flex justify-center">
                <Link href="/#companies">
                    <PlayIcon className="text-blue-950 w-40 h-40 rotate-90"/>
                </Link>
            </div>
            <div id="companies" className="flex justify-center">here goes crawled companies</div>
        </div>
    )
}

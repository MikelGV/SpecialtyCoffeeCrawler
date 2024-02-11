import Link from "next/link"
import { PlayIcon } from "@heroicons/react/16/solid"

export default function HomeS() {
    return (

        <div>
            <div className="flex justify-center m-56 mb-64">
                <h1 className="text-black font-pacifico text-6xl">Specialty Coffee Europe</h1>
            </div>
            <div className="flex justify-center">
                <Link href="/#companies">
                    <PlayIcon className="text-blue-950 w-40 h-40 rotate-90"/>
                </Link>
            </div>
            <div id="companies" className="flex-row justify-center m-48 font-pacifico">
                <div className="flex justify-center mb-20">
                    <h1 className="text-black text-4xl">Coffee Shops Crawled:</h1>
                </div>
                <div className="flex justify-center text-black">
                    <div id="nomad" className="flex-1">
                        <div className="flex justify-center bg-white h-48 w-72">
                            <div className="mt-20">
                                <h1 className="font-grotesque font-semibold h-5 w-15">NOMAD Coffee</h1>
                            </div>
                        </div>
                    </div>
                    <div className="flex-auto">Product2</div>
                    <div className="flex">Product3</div>
                </div>
            </div>
        </div>
    )
}

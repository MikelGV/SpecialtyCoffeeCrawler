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
            <div id="companies" className="flex-row justify-center content-center m-48 font-pacifico">
                <div className="flex justify-center mb-20">
                    <h1 className="text-black text-4xl">Coffee Shops Crawled:</h1>
                </div>
                <div className="flex justify-center text-black mx-48">
                    <div id="nomad" className="flex-1">
                        <div className="flex justify-center">
                            <div className="mt-20">
                                <Link href="https://nomadcoffee.es/en" target="_blank" className="font-grotesque text-3xl font-semibold h-5 w-15">NOMAD Coffee</Link>
                            </div>
                        </div>
                    </div>
                    <div className="flex-auto w-48 md:w-32">
                        <div className="flex justify-center">
                            <Link href="https://www.fiveelephant.com/" target="_blank">
                                <img src="https://www.fiveelephant.com/cdn/shop/files/01_5el_logo_cmyk_transperant-01_400x.png?v=1614295937"/>
                            </Link>
                        </div>
                    </div>
                    <div className="flex h-48">
                        <Link href="https://ariosacoffee.com/" target="_blank">
                            <img src="https://ariosacoffee.com/cdn/shop/t/5/assets/logo.svg?v=83156499127083417791689682072"/>
                        </Link>
                    </div>
                </div>
            </div>
        </div>
    )
}

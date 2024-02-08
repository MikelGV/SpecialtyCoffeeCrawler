import Link from "next/link"

export default function HomeS() {
    return (

        <div>
            <div className="flex justify-center m-56">
                <h1 className="text-black font-pacifico text-6xl">Specialty Coffee Europe</h1>
            </div>
            <div className="flex justify-center rounded-xl">
                <Link href="/#companies" className="w-0 rounded-lg h-0 border-l-[50px] border-l-transparent
                    border-t-[75px] border-t-blue-950
                    border-r-[50px] border-r-transparent"
                ></Link>
            </div>
            <div id="companies" className="flex justify-center">here goes crawled companies</div>
        </div>
    )
}

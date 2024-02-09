import Link from "next/link"
import ArrowDropDownRoundedIcon from '@mui/icons-material/ArrowDropDownRounded';
export default function HomeS() {
    return (

        <div>
            <div className="flex justify-center m-56">
                <h1 className="text-black font-pacifico text-6xl">Specialty Coffee Europe</h1>
            </div>
            <div className="flex justify-center">
                <Link href="/#companies">
                    <ArrowDropDownRoundedIcon className="w-96 h-96 border-black bg-blue-950"/>
                </Link>
            </div>
            <div id="companies" className="flex justify-center">here goes crawled companies</div>
        </div>
    )
}

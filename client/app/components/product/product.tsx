import Link from "next/link";

export default function Product({ item: string }) {
    return (
        <div id="container" className="">
            <div id="imgContainer" className="">
                <Link href={item.url}>
                     <img className="" src={item.img}/>
                </Link>
            </div>
            <div id="infoContainer" className="">
                <h1 className="">{item.name}</h1>
                <p className="">{item.price}</p>
            </div>
        </div>
    )
}

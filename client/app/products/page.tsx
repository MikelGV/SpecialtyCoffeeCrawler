import Link from "next/link";
import Layout from "../components/layout/layout";

async function getData() {
    const res = await fetch('http://localhost:8080/products')

    if (!res.ok) {
        throw new Error('Failed to fetch data')
    }

    const data = await res.json()
    return data;
}

export default async function Products() {
    const Products = await getData()
    console.log(Products)

    return (
        <Layout>
            <div className="grid grid-cols-3 gap-1 justify-evenly ml-64 mt-32">
                   {
                Products

                    ?
                    Products.map((p: any) => {
                        return (
                            <>
                                <div id="products" className="">
                                    <div id="imgContainer" className="">
                                        <Link href={p.Url} target="_blank">
                                            <img className="h-44 w-48" src={p.Img} />
                                        </Link>
                                    </div>
                                    <div id="infoContainer" className="text-black">
                                        <h1 className="">{p.Name}</h1>
                                        <p className="">{p.Price}</p>
                                    </div>
                                </div>
                            </>
                        )
                    })
                    :

                    "Getting Products ...."


                // THIS LOADING STATE WILL NOT BE VISIBLE BECAUSE SERVER LOADS THIS WHOLE PAGE
            }
            </div>
        </Layout>
    )
}

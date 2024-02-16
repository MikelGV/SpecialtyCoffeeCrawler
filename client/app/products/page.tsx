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
            <div className="grid grid-cols-3 gap-1 justify-evenly ml-64 mt-32 font-pacifico">
                   {
                Products

                    ?
                    Products.map((p: any) => {
                        return (
                            <>
                                <div id="products" className="">
                                    <div id="imgContainer" className="h-48 w-48">
                                        <Link href={p.Url} target="_blank">
                                            <img className="bg-white mb-5" src={p.Img} />
                                        </Link>
                                    </div>
                                    <div id="infoContainer" className="text-black w-48">
                                        <Link href={p.Url} target="_blank">
                                            <h1 className="text-xl mb-5">{p.Name}</h1>
                                        </Link>
                                        <p className="text-lg font-semibold mb-5">{p.Price}</p>
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

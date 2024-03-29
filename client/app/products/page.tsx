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
            <div className="max-w-xl xl:max-w-7xl xl:mx-70 xl:my-32 md:mx-60 md:my-32 sm:my-32 sm:ml-14">
                <div className="grid grid-cols-1 gap-1 justify-evenly sm:grid-cols-2 lg:grid-cols-3 font-pacifico">
                    {
                        Products

                        ?
                        Products.map((p: any) => {
                            return (
                                <>
                                    <div id="products" className="mb-10">
                                        <div id="imgContainer" className="">
                                            <Link className="group" href={p.Url} target="_blank">
                                                <img className="object-fill bg-white w-72 mb-5 rounded-xl group-hover:opacity-75" src={p.Img} />
                                            </Link>
                                        </div>
                                        <div id="infoContainer" className="text-black w-48">
                                            <Link className="group" href={p.Url} target="_blank">
                                                <h1 className="group-hover:opacity-75 text-xl mb-5">{p.Name}</h1>
                                            </Link>
                                            <p className="mb-5 text-lg font-semibold">{p.Price}</p>
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
            </div>
        </Layout>
    )
}

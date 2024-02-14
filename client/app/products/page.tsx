import Layout from "../components/layout/layout";

async function getData() {
    const res = await fetch('http://localhost:8080/products')

    if (!res.ok) {
        throw new Error('Failed to fetch data')
    }

    return res.json()
}

export default async function Products() {
    const data = await getData()

    return (
        <Layout>
            <div className="h-48 w-48">
                <h1 className="text-black">{data.name}</h1>
            </div>
        </Layout>
    )
}

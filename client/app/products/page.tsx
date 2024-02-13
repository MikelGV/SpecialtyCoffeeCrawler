import Layout from "../components/layout/layout";

async function getData() {
    const res = await fetch('localhost:8080/products')

    if (!res.ok) {
        throw new Error('Failed to fetch data')
    }

    return res.json()
}

export default async function Products() {
    const data = await getData()

    return (
        <Layout>
            <div>This is product Page {data}</div>
        </Layout>
    )
}

import Layout from "./components/layout/layout";
import HomeS from "./home/home";
export default function Home() {
  return (

      <Layout>
        <main className="flex min-h-screen flex-col justify-between">
            <HomeS/>
        </main>
      </Layout> 
  );
}

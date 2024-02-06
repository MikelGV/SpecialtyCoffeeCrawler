import Image from "next/image";
import TopBar from "./components/topBar/topBar";
import Footer from "./components/footer/footer";
export default function Home() {
  return (
    <main className="flex min-h-screen flex-col justify-between">
        <TopBar/>
        <h1 className="text-red-700">Hello world</h1>
        <Footer/>
    </main>
  );
}

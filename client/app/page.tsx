import Image from "next/image";
import TopBar from "./components/topBar/topBar";
import Footer from "./components/footer/footer";
export default function Home() {
  return (
    <main className="flex min-h-screen flex-col items-center justify-between p-24">
        <TopBar/>
        <h1 className="text-red-700">Hello world</h1>
        <Footer/>
    </main>
  );
}

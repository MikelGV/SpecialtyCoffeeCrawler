import TopBar from "./components/topBar/topBar";
import Footer from "./components/footer/footer";
import HomeS from "./pages/home/home";
export default function Home() {
  return (
    <main className="flex min-h-screen flex-col justify-between">
        <TopBar/>
        <HomeS/>
        <Footer/>
    </main>
  );
}

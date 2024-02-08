import Footer from "../footer/footer";
import TopBar from "../topBar/topBar";

export default function Layout({ children} : { children: React.ReactNode}) {
    return (
        <div className="flex min-h-screen flex-col justify-between">
            <TopBar/>
            {children}
            <Footer/>
        </div>
    )
}

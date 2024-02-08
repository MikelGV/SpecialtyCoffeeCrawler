import Footer from "../footer/footer";
import TopBar from "../topBar/topBar";

export default function Layout({ children} : { children: React.ReactNode}) {
    return (
        <div>
            <TopBar/>
            {children}
            <Footer/>
        </div>
    )
}


export default function TopBar() {
    return (
        <div className="flex-col">
            <div className="bg-yellow-950 p-0.5"></div>
            <div className="flex bg-green-950 p-5 rounded-b-xl font-pacifico">
                <div id="left" className="flex-1">     
                    <h1 className="">Home</h1>
                </div>
                <div id="right" className="">
                    <h1 className="">Products</h1>
                </div>
            </div>
        </div>
    )
}


export default function TopBar() {
    return (
        <div>
            <div className="flex-col bg-yellow-950 p-0.5"></div>
            <div className="flex-row bg-green-950 p-3 rounded-b-xl">
                <div id="left" className="flex-col">     
                    <h1 className="">Home</h1>
                </div>
                <div id="right" className="flex-col">
                    <h1 className="">Products</h1>
                </div>
            </div>
        </div>
    )
}

import Link from "next/link";

export default function Header() {
    return (
        <div className="w-full p-2 flex justify-between text-conquest-gold bg-conquest-silver" id="header">
            <ul id="navigation-section">
                <li className="cursor-pointer">
                    <Link href="/mod">Mods</Link>
                </li>
            </ul>
            <div>
                <span>I will exist someday</span>
            </div>
        </div>
    )
}
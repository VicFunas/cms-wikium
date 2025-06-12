export default function Banner({
    src,
    altText
}: {
    src: string,
    altText: string
}) {
    return <img src={src} className="w-full object-fill h-48" alt={altText} />
}
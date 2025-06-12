import "./styles.css"

export default function Feedback({
    type,
    label,
    content
}: {
    type: "success" | "error",
    label: string,
    content: string
}) {
    return (
        <div className={`mt-8 p-4 border rounded-md w-full text-center ${type}-feedback`}>
            <p className="font-semibold">{`${label}:`}</p>
            <p className="text-lg">{content}</p>
          </div>
    )
}
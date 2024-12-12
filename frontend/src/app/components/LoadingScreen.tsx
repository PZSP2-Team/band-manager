"use client";

export default function LoadingScreen() {
    return (
        <div className="min-h-[calc(100vh-74px)] flex items-center justify-center">
            <div className="flex flex-col items-center gap-4">
                <div className="w-8 h-8 border-4 border-t-cornflowerblue border-r-transparent border-b-cornflowerblue border-l-transparent rounded-full animate-spin" />
                <p className="text-customGray">Loading...</p>
            </div>
        </div>
    );
}

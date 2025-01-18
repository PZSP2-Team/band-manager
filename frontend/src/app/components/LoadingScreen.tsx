"use client";

/**
 * Loading spinner component displayed during data fetching or processing.
 * Renders a centered, animated spinner with "Loading..." text.
 *
 * Features:
 * - Fullscreen centered layout
 * - Animated circular spinner
 * - Loading text indicator
 * - Custom color scheme matching app theme
 */
export default function LoadingScreen() {
  return (
    <div className="h-full flex items-center justify-center">
      <div className="flex flex-col items-center gap-4">
        <div className="w-8 h-8 border-4 border-t-cornflowerblue border-r-transparent border-b-cornflowerblue border-l-transparent rounded-full animate-spin" />
        <p className="text-customGray">Loading...</p>
      </div>
    </div>
  );
}

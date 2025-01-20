import { NextRequest, NextResponse } from "next/server";

/**
 * Proxy endpoint for notesheet creation.
 * Forwards POST request to backend service and handles response.
 *
 * Endpoint: POST /api/track/notesheet/create
 *
 * @description
 * Acts as a proxy between frontend and backend services:
 * - Forwards incoming request (body and headers) to backend
 * - Handles backend response and error states
 * - Maintains same status codes from backend
 *
 * Error Handling:
 * - Returns backend error response if request fails
 * - Returns 500 status for unexpected errors
 */
export async function POST(req: NextRequest) {
  try {
    const response = await fetch(
      `http://${process.env.NEXT_PUBLIC_BACKEND_HOST}:${process.env.NEXT_PUBLIC_BACKEND_PORT}/api/track/notesheet/create/`,
      {
        method: "POST",
        body: req.body,
        headers: req.headers,
        // @ts-expect-error - ignore the error
        duplex: "half",
      },
    );

    if (!response.ok) {
      const errorText = await response.text();
      console.error("Backend error:", errorText);
      return NextResponse.json(
        { error: "Backend request failed" },
        { status: response.status },
      );
    }

    const data = await response.json();
    return NextResponse.json(data);
  } catch (error) {
    console.error("Error:", error);
    return NextResponse.json(
      { error: "Internal server error" },
      { status: 500 },
    );
  }
}

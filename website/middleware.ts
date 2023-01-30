import { NextRequest, NextResponse } from "next/server";

export function middleware(request: NextRequest) {
  const cookie = request.cookies.get("jwt")

  if (cookie && request.nextUrl.pathname.startsWith("/login")){
    return NextResponse.redirect(new URL("/", request.url))
  }

  /* if (!cookie && request.nextUrl.pathname.startsWith("/upload")){
    return NextResponse.redirect(new URL("/login", request.url))
  } */

  return NextResponse.next()
}

export const config = { 
  matcher: ['/login', '/upload']
}

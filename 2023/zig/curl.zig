// --library curl --library c $(pkg-config --cflags libcurl)`
const std = @import("std");
const String = @import("string.zig").String;
const cURL = @cImport({
    @cInclude("curl/curl.h");
});

pub fn curl(day: []const u8, session: []const u8) !void {
    var arena_state = std.heap.ArenaAllocator.init(std.heap.c_allocator);
    defer arena_state.deinit();

    const allocator = arena_state.allocator();

    // global curl init, or fail
    if (cURL.curl_global_init(cURL.CURL_GLOBAL_ALL) != cURL.CURLE_OK)
        return error.CURLGlobalInitFailed;
    defer cURL.curl_global_cleanup();

    // curl easy handle init, or fail
    const handle = cURL.curl_easy_init() orelse return error.CURLHandleInitFailed;
    defer cURL.curl_easy_cleanup(handle);

    var response_buffer = std.ArrayList(u8).init(allocator);

    // superfluous when using an arena allocator, but
    // important if the allocator implementation changes
    defer response_buffer.deinit();

    // setup curl options
    var url = String.init(allocator);
    defer url.deinit();
    try url.concat("https://adventofcode.com/2023/day/");
    try url.concat(day);
    try url.concat("/input");

    if (cURL.curl_easy_setopt(handle, cURL.CURLOPT_URL, url.str()) != cURL.CURLE_OK)
        return error.CouldNotSetURL;

    const cookie = String.init(allocator);
    defer cookie.deinit();
    try cookie.concat("session=");
    try cookie.concat(session);
    try cookie.concat(";");

    // setup session cookie
    if (cURL.curl_easy_setopt(handle, cURL.CURLOPT_COOKIE, cookie.str()) != cURL.CURLE_OK)
        return error.CouldNotSetSessionCookie;

    // set write function callbacks
    if (cURL.curl_easy_setopt(handle, cURL.CURLOPT_WRITEFUNCTION, writeToArrayListCallback) != cURL.CURLE_OK)
        return error.CouldNotSetWriteCallback;
    if (cURL.curl_easy_setopt(handle, cURL.CURLOPT_WRITEDATA, &response_buffer) != cURL.CURLE_OK)
        return error.CouldNotSetWriteCallback;

    // perform
    if (cURL.curl_easy_perform(handle) != cURL.CURLE_OK)
        return error.FailedToPerformRequest;
}

fn writeToArrayListCallback(data: *anyopaque, size: c_uint, nmemb: c_uint, user_data: *anyopaque) callconv(.C) c_uint {
    var buffer: *std.ArrayList(u8) = @alignCast(@ptrCast(user_data));
    var typed_data: [*]u8 = @ptrCast(data);
    buffer.appendSlice(typed_data[0 .. nmemb * size]) catch return 0;
    return nmemb * size;
}

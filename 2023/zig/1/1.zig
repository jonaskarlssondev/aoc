const std = @import("std");

pub fn main() !void {
    var arena = std.heap.ArenaAllocator.init(std.heap.page_allocator);
    defer arena.deinit();

    const file = try std.fs.cwd().openFile("./input.txt", .{});
    defer file.close();

    const content = try file.reader().readAllAlloc(&arena.allocator(), 10 * 1024 * 1024);
    defer arena.allocator().free(content);

    std.debug.print(content);
}

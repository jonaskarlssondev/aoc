const std = @import("std");

pub const String = struct {
    buffer: ?[]u8,
    allocator: std.mem.Allocator,
    size: usize,

    pub const Error = error{
        OutOfMemory,
    };

    pub fn init(allocator: std.mem.Allocator) String {
        return .{
            .buffer = null,
            .allocator = allocator,
            .size = 0,
        };
    }

    pub fn deinit(self: *String) void {
        if (self.buffer) |buf| self.allocator.free(buf);
    }

    pub fn allocate(self: *String, bytes: usize) Error!void {
        if (self.buffer) |buf| {
            self.buffer = self.allocator.realloc(buf, bytes) catch {
                return Error.OutOfMemory;
            };
        } else {
            self.buffer = self.allocator.alloc(u8, bytes) catch {
                return Error.OutOfMemory;
            };
        }
    }

    pub fn insert(self: *String, inp: []const u8) Error!void {
        if (self.buffer) |buf| {
            if (self.size + inp.len > buf.len) {
                try self.allocate(2 * (self.size + inp.len));
            }
        } else {
            try self.allocate(2 * inp.len);
        }

        const buf = self.buffer.?;

        var i: usize = 0;
        while (i < inp.len) : (i += 1) {
            buf[self.size + i] = inp[i];
        }

        self.size += inp.len;
    }

    pub fn concat(self: *String, inp: []const u8) Error!void {
        try self.insert(inp);
    }

    pub fn str(self: String) []const u8 {
        if (self.buffer) |buf| {
            return buf[0..self.size];
        }

        return "";
    }
};

import XCTest
import SwiftTreeSitter
import TreeSitterHoatzin

final class TreeSitterHoatzinTests: XCTestCase {
    func testCanLoadGrammar() throws {
        let parser = Parser()
        let language = Language(language: tree_sitter_hoatzin())
        XCTAssertNoThrow(try parser.setLanguage(language),
                         "Error loading Hoatzin grammar")
    }
}

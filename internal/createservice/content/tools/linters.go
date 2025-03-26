package tools

const LintersContent = `# This file contains all available configuration options
# with their default values (in comments).
#
# This file is not a configuration example,
# it contains the exhaustive configuration with explanations of the options.

version: "2"

linters:
  # Default set of linters.
  # The value can be: standard, all, none, or fast.
  # Default: standard
  default: all

  # Enable specific linter.
  # https://golangci-lint.run/usage/linters/#enabled-by-default
  enable:
    - asasalint
    - asciicheck
    - bidichk
    - bodyclose
    - canonicalheader
    - containedctx
    - contextcheck
    - copyloopvar
    - cyclop
    - decorder
    - depguard
    - dogsled
    - dupl
    - dupword
    - durationcheck
    - err113
    - errcheck
    - errchkjson
    - errname
    - errorlint
    - exhaustive
    - exhaustruct
    - exptostd
    - fatcontext
    - forbidigo
    - forcetypeassert
    - funlen
    - ginkgolinter
    - gocheckcompilerdirectives
    - gochecknoglobals
    - gochecknoinits
    - gochecksumtype
    - gocognit
    - goconst
    - gocritic
    - gocyclo
    - godot
    - godox
    - goheader
    - gomoddirectives
    - gomodguard
    - goprintffuncname
    - gosec
    - gosmopolitan
    - govet
    - grouper
    - iface
    - importas
    - inamedparam
    - ineffassign
    - interfacebloat
    - intrange
    - ireturn
    - lll
    - loggercheck
    - maintidx
    - makezero
    - mirror
    - misspell
    - mnd
    - musttag
    - nakedret
    - nestif
    - nilerr
    - nilnesserr
    - nilnil
    - nlreturn
    - noctx
    - nolintlint
    - nonamedreturns
    - nosprintfhostport
    - paralleltest
    - perfsprint
    - prealloc
    - predeclared
    - promlinter
    - protogetter
    - reassign
    - recvcheck
    - revive
    - rowserrcheck
    - sloglint
    - spancheck
    - sqlclosecheck
    - staticcheck
    - tagalign
    - tagliatelle
    - testableexamples
    - testifylint
    - testpackage
    - thelper
    - tparallel
    - unconvert
    - unparam
    - unused
    - usestdlibvars
    - usetesting
    - varnamelen
    - wastedassign
    - whitespace
    - wrapcheck
    - wsl
    - zerologlint

  # Disable specific linter
  # https://golangci-lint.run/usage/linters/#disabled-by-default
  disable:
    - asasalint
    - asciicheck
    - bidichk
    - bodyclose
    - canonicalheader
    - containedctx
    - contextcheck
    - copyloopvar
    - cyclop
    - decorder
    - depguard
    - dogsled
    - dupl
    - dupword
    - durationcheck
    - err113
    - errcheck
    - errchkjson
    - errname
    - errorlint
    - exhaustive
    - exhaustruct
    - exptostd
    - fatcontext
    - forbidigo
    - forcetypeassert
    - funlen
    - ginkgolinter
    - gocheckcompilerdirectives
    - gochecknoglobals
    - gochecknoinits
    - gochecksumtype
    - gocognit
    - goconst
    - gocritic
    - gocyclo
    - godot
    - godox
    - goheader
    - gomoddirectives
    - gomodguard
    - goprintffuncname
    - gosec
    - gosmopolitan
    - govet
    - grouper
    - iface
    - importas
    - inamedparam
    - ineffassign
    - interfacebloat
    - intrange
    - ireturn
    - lll
    - loggercheck
    - maintidx
    - makezero
    - mirror
    - misspell
    - mnd
    - musttag
    - nakedret
    - nestif
    - nilerr
    - nilnesserr
    - nilnil
    - nlreturn
    - noctx
    - nolintlint
    - nonamedreturns
    - nosprintfhostport
    - paralleltest
    - perfsprint
    - prealloc
    - predeclared
    - promlinter
    - protogetter
    - reassign
    - recvcheck
    - revive
    - rowserrcheck
    - sloglint
    - spancheck
    - sqlclosecheck
    - staticcheck
    - tagalign
    - tagliatelle
    - testableexamples
    - testifylint
    - testpackage
    - thelper
    - tparallel
    - unconvert
    - unparam
    - unused
    - usestdlibvars
    - usetesting
    - varnamelen
    - wastedassign
    - whitespace
    - wrapcheck
    - wsl
    - zerologlint

  # All available settings of specific linters.
  settings:
    asasalint:
      # To specify a set of function names to exclude.
      # The values are merged with the builtin exclusions.
      # The builtin exclusions can be disabled by setting use-builtin-exclusions to false.
      # Default: ["^(fmt|log|logger|t|)\.(Print|Fprint|Sprint|Fatal|Panic|Error|Warn|Warning|Info|Debug|Log)(|f|ln)$"]
      exclude:
        - Append
        - \.Wrapf
      # To enable/disable the asasalint builtin exclusions of function names.
      # See the default value of exclude to get the builtin exclusions.
      # Default: true
      use-builtin-exclusions: false

    bidichk:
      # The following configurations check for all mentioned invisible Unicode runes.
      # All runes are enabled by default.
      left-to-right-embedding: false
      right-to-left-embedding: false
      pop-directional-formatting: false
      left-to-right-override: false
      right-to-left-override: false
      left-to-right-isolate: false
      right-to-left-isolate: false
      first-strong-isolate: false
      pop-directional-isolate: false

    copyloopvar:
      # Check all assigning the loop variable to another variable.
      # Default: false
      check-alias: true

    cyclop:
      # The maximal code complexity to report.
      # Default: 10
      max-complexity: 10
      # The maximal average package complexity.
      # If it's higher than 0.0 (float) the check is enabled
      # Default: 0.0
      package-average: 0.5

    decorder:
      # Required order of type, const, var and func declarations inside a file.
      # Default: types before constants before variables before functions.
      dec-order:
        - type
        - const
        - var
        - func

      # If true, underscore vars (vars with "_" as the name) will be ignored at all checks
      # Default: false (underscore vars are not ignored)
      ignore-underscore-vars: false

      # If true, order of declarations is not checked at all.
      # Default: true (disabled)
      disable-dec-order-check: false

      # If true, init func can be anywhere in file (does not have to be declared before all other functions).
      # Default: true (disabled)
      disable-init-func-first-check: false

      # If true, multiple global type, const and var declarations are allowed.
      # Default: true (disabled)
      disable-dec-num-check: false

      # If true, type declarations will be ignored for dec num check
      # Default: false (type statements are not ignored)
      disable-type-dec-num-check: false

      # If true, const declarations will be ignored for dec num check
      # Default: false (const statements are not ignored)
      disable-const-dec-num-check: false

      # If true, var declarations will be ignored for dec num check
      # Default: false (var statements are not ignored)
      disable-var-dec-num-check: false

    depguard:
      # Rules to apply.
      #
      # Variables:
      # - File Variables
      #   Use an exclamation mark ! to negate a variable.
      #   Example: !$test matches any file that is not a go test file.
      #
      #   $all - matches all go files
      #   $test - matches all go test files
      #
      # - Package Variables
      #
      #   $gostd - matches all of go's standard library (Pulled from GOROOT)
      #
      # Default (applies if no custom rules are defined): Only allow $gostd in all files.
      rules:
        # Name of a rule.
        main:
          # Defines package matching behavior. Available modes:
          # - original: allowed if it doesn't match the deny list and either matches the allow list or the allow list is empty.
          # - strict: allowed only if it matches the allow list and either doesn't match the deny list or the allow rule is more specific (longer) than the deny rule.
          # - lax: allowed if it doesn't match the deny list or the allow rule is more specific (longer) than the deny rule.
          # Default: "original"
          list-mode: lax
          # List of file globs that will match this list of settings to compare against.
          # By default, if a path is relative, it is relative to the directory where the golangci-lint command is executed.
          # The placeholder '${base-path}' is substituted with a path relative to the mode defined with run.relative-path-mode.
          # Default: $all
          files:
            - "!**/*_a _file.go"
          # List of allowed packages.
          # Entries can be a variable (starting with $), a string prefix, or an exact match (if ending with $).
          # Default: []
          allow:
            - $gostd
            - github.com/OpenPeeDeeP
          # List of packages that are not allowed.
          # Entries can be a variable (starting with $), a string prefix, or an exact match (if ending with $).
          # Default: []
          deny:
            - pkg: "math/rand$"
              desc: use math/rand/v2
            - pkg: "github.com/sirupsen/logrus"
              desc: not allowed
            - pkg: "github.com/pkg/errors"
              desc: Should be replaced by standard lib errors package

    dogsled:
      # Checks assignments with too many blank identifiers.
      # Default: 2
      max-blank-identifiers: 3

    dupl:
      # Tokens count to trigger issue.
      # Default: 150
      threshold: 100

    dupword:
      # Keywords for detecting duplicate words.
      # If this list is not empty, only the words defined in this list will be detected.
      # Default: []
      keywords:
        - "the"
        - "and"
        - "a"
      # Keywords used to ignore detection.
      # Default: []
      ignore:
        - "0C0C"

    errcheck:
      # Report about not checking of errors in type assertions: a := b.(MyStruct).
      # Such cases aren't reported by default.
      # Default: false
      check-type-assertions: true

      # report about assignment of errors to blank identifier: num, _ := strconv.Atoi(numStr).
      # Such cases aren't reported by default.
      # Default: false
      check-blank: true

      # To disable the errcheck built-in exclude list.
      # See -excludeonly option in https://github.com/kisielk/errcheck#excluding-functions for details.
      # Default: false
      disable-default-exclusions: true

      # List of functions to exclude from checking, where each entry is a single function to exclude.
      # See https://github.com/kisielk/errcheck#excluding-functions for details.
      exclude-functions:
        - io/ioutil.ReadFile
        - io.Copy(*bytes.Buffer)
        - io.Copy(os.Stdout)

    errchkjson:
      # With check-error-free-encoding set to true, errchkjson does warn about errors
      # from json encoding functions that are safe to be ignored,
      # because they are not possible to happen.
      #
      # if check-error-free-encoding is set to true and errcheck linter is enabled,
      # it is recommended to add the following exceptions to prevent from false positives:
      #
      #     linters-settings:
      #       errcheck:
      #         exclude-functions:
      #           - encoding/json.Marshal
      #           - encoding/json.MarshalIndent
      #
      # Default: false
      check-error-free-encoding: true

      # Issue on struct encoding that doesn't have exported fields.
      # Default: false
      report-no-exported: false

    errorlint:
      # Check whether fmt.Errorf uses the %w verb for formatting errors.
      # See the https://github.com/polyfloyd/go-errorlint for caveats.
      # Default: true
      errorf: false
      # Permit more than 1 %w verb, valid per Go 1.20 (Requires errorf:true)
      # Default: true
      errorf-multi: false
      # Check for plain type assertions and type switches.
      # Default: true
      asserts: false
      # Check for plain error comparisons.
      # Default: true
      comparison: false
      # Allowed errors.
      # Default: []
      allowed-errors:
        - err: "io.EOF"
          fun: "example.com/pkg.Read"
      # Allowed error "wildcards".
      # Default: []
      allowed-errors-wildcard:
        - err: "example.com/pkg.ErrMagic"
          fun: "example.com/pkg.Magic"

    exhaustive:
      # Program elements to check for exhaustiveness.
      # Default: [ switch ]
      check:
        - switch
        - map
      # Presence of "default" case in switch statements satisfies exhaustiveness,
      # even if all enum members are not listed.
      # Default: false
      default-signifies-exhaustive: true
      # Enum members matching the supplied regex do not have to be listed in
      # switch statements to satisfy exhaustiveness.
      # Default: ""
      ignore-enum-members: "Example.+"
      # Enum types matching the supplied regex do not have to be listed in
      # switch statements to satisfy exhaustiveness.
      # Default: ""
      ignore-enum-types: "Example.+"
      # Consider enums only in package scopes, not in inner scopes.
      # Default: false
      package-scope-only: true
      # Only run exhaustive check on switches with "//exhaustive:enforce" comment.
      # Default: false
      explicit-exhaustive-switch: true
      # Only run exhaustive check on map literals with "//exhaustive:enforce" comment.
      # Default: false
      explicit-exhaustive-map: true
      # Switch statement requires default case even if exhaustive.
      # Default: false
      default-case-required: true

    exhaustruct:
      # List of regular expressions to match struct packages and their names.
      # Regular expressions must match complete canonical struct package/name/structname.
      # If this list is empty, all structs are tested.
      # Default: []
      include:
        - '.+\.Test'
        - 'example\.com/package\.ExampleStruct[\d]{1,2}'
      # List of regular expressions to exclude struct packages and their names from checks.
      # Regular expressions must match complete canonical struct package/name/structname.
      # Default: []
      exclude:
        - '.+/cobra\.Command$'

    fatcontext:
      # Check for potential fat contexts in struct pointers.
      # May generate false positives.
      # Default: false
      check-struct-pointers: true

    forbidigo:
      # Forbid the following identifiers (list of regexp).
      # Default: ["^(fmt\\.Print(|f|ln)|print|println)$"]
      forbid:
        # Built-in bootstrapping functions.
        - pattern: ^print(ln)?$
        # Optional message that gets included in error reports.
        - pattern: ^fmt\.Print.*$
          msg: Do not commit print statements.
        # Alternatively, put messages at the end of the regex, surrounded by (# )?
        # Escape any special characters. Those messages get included in error reports.
        - pattern: 'fmt\.Print.*(# Do not commit print statements\.)?'
        # Forbid spew Dump, whether it is called as function or method.
        # Depends on analyze-types below.
        - pattern: ^spew\.(ConfigState\.)?Dump$
        # The package name might be ambiguous.
        # The full import path can be used as additional criteria.
        # Depends on analyze-types below.
        - pattern: ^v1.Dump$
          pkg: ^example.com/pkg/api/v1$
      # Exclude godoc examples from forbidigo checks.
      # Default: true
      exclude-godoc-examples: false
      # Instead of matching the literal source code,
      # use type information to replace expressions with strings that contain the package name
      # and (for methods and fields) the type name.
      # This makes it possible to handle import renaming and forbid struct fields and methods.
      # Default: false
      analyze-types: true

    funlen:
      # Checks the number of lines in a function.
      # If lower than 0, disable the check.
      # Default: 60
      lines: -1
      # Checks the number of statements in a function.
      # If lower than 0, disable the check.
      # Default: 40
      statements: -1
      # Ignore comments when counting lines.
      # Default: true
      ignore-comments: false

    ginkgolinter:
      # Suppress the wrong length assertion warning.
      # Default: false
      suppress-len-assertion: true

      # Suppress the wrong nil assertion warning.
      # Default: false
      suppress-nil-assertion: true

      # Suppress the wrong error assertion warning.
      # Default: false
      suppress-err-assertion: true

      # Suppress the wrong comparison assertion warning.
      # Default: false
      suppress-compare-assertion: true

      # Suppress the function all in async assertion warning.
      # Default: false
      suppress-async-assertion: true

      # Suppress warning for comparing values from different types, like int32 and uint32
      # Default: false
      suppress-type-compare-assertion: true

      # Trigger warning for ginkgo focus containers like FDescribe, FContext, FWhen or FIt
      # Default: false
      forbid-focus-container: true

      # Don't trigger warnings for HaveLen(0)
      # Default: false
      allow-havelen-zero: true

      # Force using Expect with To, ToNot or NotTo.
      # Reject using Expect with Should or ShouldNot.
      # Default: false
      force-expect-to: true

      # Best effort validation of async intervals (timeout and polling).
      # Ignored the suppress-async-assertion is true.
      # Default: false
      validate-async-intervals: true

      # Trigger a warning for variable assignments in ginkgo containers like Describe, Context and When, instead of in BeforeEach().
      # Default: false
      forbid-spec-pollution: true

      # Force using the Succeed matcher for error functions, and the HaveOccurred matcher for non-function error values.
      # Default: false
      force-succeed: true

    gochecksumtype:
      # Presence of default case in switch statements satisfies exhaustiveness, if all members are not listed.
      # Default: true
      default-signifies-exhaustive: false
      # Include shared interfaces in the exhaustiviness check.
      # Default: false
      include-shared-interfaces: true

    gocognit:
      # Minimal code complexity to report.
      # Default: 30 (but we recommend 10-20)
      min-complexity: 10

    goconst:
      # Minimal length of string constant.
      # Default: 3
      min-len: 2
      # Minimum occurrences of constant string count to trigger issue.
      # Default: 3
      min-occurrences: 2
      # Look for existing constants matching the values.
      # Default: true
      match-constant: false
      # Search also for duplicated numbers.
      # Default: false
      numbers: true
      # Minimum value, only works with goconst.numbers
      # Default: 3
      min: 2
      # Maximum value, only works with goconst.numbers
      # Default: 3
      max: 2
      # Ignore when constant is not used as function argument.
      # Default: true
      ignore-calls: false
      # Exclude strings matching the given regular expression.
      # Default: ""
      ignore-strings: 'foo.+'

    gocritic:
      # Disable all checks.
      # Default: false
      disable-all: true
      # Which checks should be enabled in addition to default checks; can't be combined with 'disabled-checks'.
      # By default, list of stable checks is used (https://go-critic.com/overview#checks-overview):
      #   appendAssign, argOrder, assignOp, badCall, badCond, captLocal, caseOrder, codegenComment, commentFormatting,
      #   defaultCaseOrder, deprecatedComment, dupArg, dupBranchBody, dupCase, dupSubExpr, elseif, exitAfterDefer,
      #   flagDeref, flagName, ifElseChain, mapKey, newDeref, offBy1, regexpMust, singleCaseSwitch, sloppyLen,
      #   sloppyTypeAssert, switchTrue, typeSwitchVar, underef, unlambda, unslice, valSwap, wrapperFunc
      # To see which checks are enabled run GL_DEBUG=gocritic golangci-lint run --enable=gocritic.
      enabled-checks:
        # Detects suspicious append result assignments.
        # https://go-critic.com/overview.html#appendassign
        - appendAssign
        # Detects append chains to the same slice that can be done in a single append call.
        # https://go-critic.com/overview.html#appendcombine
        - appendCombine
        # Detects suspicious arguments order.
        # https://go-critic.com/overview.html#argorder
        - argOrder
        # Detects assignments that can be simplified by using assignment operators.
        # https://go-critic.com/overview.html#assignop
        - assignOp
        # Detects suspicious function calls.
        # https://go-critic.com/overview.html#badcall
        - badCall
        # Detects suspicious condition expressions.
        # https://go-critic.com/overview.html#badcond
        - badCond
        # Detects suspicious mutex lock/unlock operations.
        # https://go-critic.com/overview.html#badlock
        - badLock
        # Detects suspicious regexp patterns.
        # https://go-critic.com/overview.html#badregexp
        - badRegexp
        # Detects bad usage of sort package.
        # https://go-critic.com/overview.html#badsorting
        - badSorting
        # Detects bad usage of sync.OnceFunc.
        # https://go-critic.com/overview.html#badsynconcefunc
        - badSyncOnceFunc
        # Detects bool expressions that can be simplified.
        # https://go-critic.com/overview.html#boolexprsimplify
        - boolExprSimplify
        # Detects when predeclared identifiers are shadowed in assignments.
        # https://go-critic.com/overview.html#builtinshadow
        - builtinShadow
        # Detects top-level declarations that shadow the predeclared identifiers.
        # https://go-critic.com/overview.html#builtinshadowdecl
        - builtinShadowDecl
        # Detects capitalized names for local variables.
        # https://go-critic.com/overview.html#captlocal
        - captLocal
        # Detects erroneous case order inside switch statements.
        # https://go-critic.com/overview.html#caseorder
        - caseOrder
        # Detects malformed 'code generated' file comments.
        # https://go-critic.com/overview.html#codegencomment
        - codegenComment
        # Detects comments with non-idiomatic formatting.
        # https://go-critic.com/overview.html#commentformatting
        - commentFormatting
        # Detects commented-out code inside function bodies.
        # https://go-critic.com/overview.html#commentedoutcode
        - commentedOutCode
        # Detects commented-out imports.
        # https://go-critic.com/overview.html#commentedoutimport
        - commentedOutImport
        # Detects when default case in switch isn't on 1st or last position.
        # https://go-critic.com/overview.html#defaultcaseorder
        - defaultCaseOrder
        # Detects loops inside functions that use defer.
        # https://go-critic.com/overview.html#deferinloop
        - deferInLoop
        # Detects deferred function literals that can be simplified.
        # https://go-critic.com/overview.html#deferunlambda
        - deferUnlambda
        # Detects malformed 'deprecated' doc-comments.
        # https://go-critic.com/overview.html#deprecatedcomment
        - deprecatedComment
        # Detects comments that silence go lint complaints about doc-comment.
        # https://go-critic.com/overview.html#docstub
        - docStub
        # Detects suspicious duplicated arguments.
        # https://go-critic.com/overview.html#duparg
        - dupArg
        # Detects duplicated branch bodies inside conditional statements.
        # https://go-critic.com/overview.html#dupbranchbody
        - dupBranchBody
        # Detects duplicated case clauses inside switch or select statements.
        # https://go-critic.com/overview.html#dupcase
        - dupCase
        # Detects multiple imports of the same package under different aliases.
        # https://go-critic.com/overview.html#dupimport
        - dupImport
        # Detects suspicious duplicated sub-expressions.
        # https://go-critic.com/overview.html#dupsubexpr
        - dupSubExpr
        # Detects suspicious formatting strings usage.
        # https://go-critic.com/overview.html#dynamicfmtstring
        - dynamicFmtString
        # Detects else with nested if statement that can be replaced with else-if.
        # https://go-critic.com/overview.html#elseif
        - elseif
        # Detects suspicious empty declarations blocks.
        # https://go-critic.com/overview.html#emptydecl
        - emptyDecl
        # Detects fallthrough that can be avoided by using multi case values.
        # https://go-critic.com/overview.html#emptyfallthrough
        - emptyFallthrough
        # Detects empty string checks that can be written more idiomatically.
        # https://go-critic.com/overview.html#emptystringtest
        - emptyStringTest
        # Detects unoptimal strings/bytes case-insensitive comparison.
        # https://go-critic.com/overview.html#equalfold
        - equalFold
        # Detects unwanted dependencies on the evaluation order.
        # https://go-critic.com/overview.html#evalorder
        - evalOrder
        # Detects calls to exit/fatal inside functions that use defer.
        # https://go-critic.com/overview.html#exitafterdefer
        - exitAfterDefer
        # Detects exposed methods from sync.Mutex and sync.RWMutex.
        # https://go-critic.com/overview.html#exposedsyncmutex
        - exposedSyncMutex
        # Detects suspicious reassignment of error from another package.
        # https://go-critic.com/overview.html#externalerrorreassign
        - externalErrorReassign
        # Detects problems in filepath.Join() function calls.
        # https://go-critic.com/overview.html#filepathjoin
        - filepathJoin
        # Detects immediate dereferencing of flag package pointers.
        # https://go-critic.com/overview.html#flagderef
        - flagDeref
        # Detects suspicious flag names.
        # https://go-critic.com/overview.html#flagname
        - flagName
        # Detects hex literals that have mixed case letter digits.
        # https://go-critic.com/overview.html#hexliteral
        - hexLiteral
        # Detects nil usages in http.NewRequest calls, suggesting http.NoBody as an alternative.
        # https://go-critic.com/overview.html#httpnobody
        - httpNoBody
        # Detects params that incur excessive amount of copying.
        # https://go-critic.com/overview.html#hugeparam
        - hugeParam
        # Detects repeated if-else statements and suggests to replace them with switch statement.
        # https://go-critic.com/overview.html#ifelsechain
        - ifElseChain
        # Detects when imported package names shadowed in the assignments.
        # https://go-critic.com/overview.html#importshadow
        - importShadow
        # Detects strings.Index calls that may cause unwanted allocs.
        # https://go-critic.com/overview.html#indexalloc
        - indexAlloc
        # Detects non-assignment statements inside if/switch init clause.
        # https://go-critic.com/overview.html#initclause
        - initClause
        # Detects suspicious map literal keys.
        # https://go-critic.com/overview.html#mapkey
        - mapKey
        # Detects method expression call that can be replaced with a method call.
        # https://go-critic.com/overview.html#methodexprcall
        - methodExprCall
        # Finds where nesting level could be reduced.
        # https://go-critic.com/overview.html#nestingreduce
        - nestingReduce
        # Detects immediate dereferencing of new expressions.
        # https://go-critic.com/overview.html#newderef
        - newDeref
        # Detects return statements those results evaluate to nil.
        # https://go-critic.com/overview.html#nilvalreturn
        - nilValReturn
        # Detects old-style octal literals.
        # https://go-critic.com/overview.html#octalliteral
        - octalLiteral
        # Detects various off-by-one kind of errors.
        # https://go-critic.com/overview.html#offby1
        - offBy1
        # Detects if function parameters could be combined by type and suggest the way to do it.
        # https://go-critic.com/overview.html#paramtypecombine
        - paramTypeCombine
        # Detects expressions like []rune(s)[0] that may cause unwanted rune slice allocation.
        # https://go-critic.com/overview.html#preferdecoderune
        - preferDecodeRune
        # Detects concatenation with os.PathSeparator which can be replaced with filepath.Join.
        # https://go-critic.com/overview.html#preferfilepathjoin
        - preferFilepathJoin
        # Detects fmt.Sprint(f/ln) calls which can be replaced with fmt.Fprint(f/ln).
        # https://go-critic.com/overview.html#preferfprint
        - preferFprint
        # Detects w.Write or io.WriteString calls which can be replaced with w.WriteString.
        # https://go-critic.com/overview.html#preferstringwriter
        - preferStringWriter
        # Detects WriteRune calls with rune literal argument that is single byte and reports to use WriteByte instead.
        # https://go-critic.com/overview.html#preferwritebyte
        - preferWriteByte
        # Detects input and output parameters that have a type of pointer to referential type.
        # https://go-critic.com/overview.html#ptrtorefparam
        - ptrToRefParam
        # Detects append all its data while range it.
        # https://go-critic.com/overview.html#rangeappendall
        - rangeAppendAll
        # Detects expensive copies of for loop range expressions.
        # https://go-critic.com/overview.html#rangeexprcopy
        - rangeExprCopy
        # Detects loops that copy big objects during each iteration.
        # https://go-critic.com/overview.html#rangevalcopy
        - rangeValCopy
        # Detects redundant fmt.Sprint calls.
        # https://go-critic.com/overview.html#redundantsprint
        - redundantSprint
        # Detects regexp.Compile* that can be replaced with regexp.MustCompile*.
        # https://go-critic.com/overview.html#regexpmust
        - regexpMust
        # Detects suspicious regexp patterns.
        # https://go-critic.com/overview.html#regexppattern
        - regexpPattern
        # Detects regexp patterns that can be simplified.
        # https://go-critic.com/overview.html#regexpsimplify
        - regexpSimplify
        # Detects suspicious http.Error call without following return.
        # https://go-critic.com/overview.html#returnafterhttperror
        - returnAfterHttpError
        # Runs user-defined rules using ruleguard linter.
        # https://go-critic.com/overview.html#ruleguard
        - ruleguard
        # Detects switch statements that could be better written as if statement.
        # https://go-critic.com/overview.html#singlecaseswitch
        - singleCaseSwitch
        # Detects slice clear loops, suggests an idiom that is recognized by the Go compiler.
        # https://go-critic.com/overview.html#sliceclear
        - sliceClear
        # Detects usage of len when result is obvious or doesn't make sense.
        # https://go-critic.com/overview.html#sloppylen
        - sloppyLen
        # Detects suspicious/confusing re-assignments.
        # https://go-critic.com/overview.html#sloppyreassign
        - sloppyReassign
        # Detects redundant type assertions.
        # https://go-critic.com/overview.html#sloppytypeassert
        - sloppyTypeAssert
        # Detects suspicious sort.Slice calls.
        # https://go-critic.com/overview.html#sortslice
        - sortSlice
        # Detects "%s" formatting directives that can be replaced with %q.
        # https://go-critic.com/overview.html#sprintfquotedstring
        - sprintfQuotedString
        # Detects issue in Query() and Exec() calls.
        # https://go-critic.com/overview.html#sqlquery
        - sqlQuery
        # Detects string concat operations that can be simplified.
        # https://go-critic.com/overview.html#stringconcatsimplify
        - stringConcatSimplify
        # Detects redundant conversions between string and []byte.
        # https://go-critic.com/overview.html#stringxbytes
        - stringXbytes
        # Detects strings.Compare usage.
        # https://go-critic.com/overview.html#stringscompare
        - stringsCompare
        # Detects switch-over-bool statements that use explicit true tag value.
        # https://go-critic.com/overview.html#switchtrue
        - switchTrue
        # Detects sync.Map load+delete operations that can be replaced with LoadAndDelete.
        # https://go-critic.com/overview.html#syncmaploadanddelete
        - syncMapLoadAndDelete
        # Detects manual conversion to milli- or microseconds.
        # https://go-critic.com/overview.html#timeexprsimplify
        - timeExprSimplify
        # Detects TODO comments without detail/assignee.
        # https://go-critic.com/overview.html#todocommentwithoutdetail
        - todoCommentWithoutDetail
        # Detects function with too many results.
        # https://go-critic.com/overview.html#toomanyresultschecker
        - tooManyResultsChecker
        # Detects potential truncation issues when comparing ints of different sizes.
        # https://go-critic.com/overview.html#truncatecmp
        - truncateCmp
        # Detects repeated type assertions and suggests to replace them with type switch statement.
        # https://go-critic.com/overview.html#typeassertchain
        - typeAssertChain
        # Detects method declarations preceding the type definition itself.
        # https://go-critic.com/overview.html#typedeffirst
        - typeDefFirst
        # Detects type switches that can benefit from type guard clause with variable.
        # https://go-critic.com/overview.html#typeswitchvar
        - typeSwitchVar
        # Detects unneeded parenthesis inside type expressions and suggests to remove them.
        # https://go-critic.com/overview.html#typeunparen
        - typeUnparen
        # Detects unchecked errors in if statements.
        # https://go-critic.com/overview.html#uncheckedinlineerr
        - uncheckedInlineErr
        # Detects dereference expressions that can be omitted.
        # https://go-critic.com/overview.html#underef
        - underef
        # Detects redundant statement labels.
        # https://go-critic.com/overview.html#unlabelstmt
        - unlabelStmt
        # Detects function literals that can be simplified.
        # https://go-critic.com/overview.html#unlambda
        - unlambda
        # Detects unnamed results that may benefit from names.
        # https://go-critic.com/overview.html#unnamedresult
        - unnamedResult
        # Detects unnecessary braced statement blocks.
        # https://go-critic.com/overview.html#unnecessaryblock
        - unnecessaryBlock
        # Detects redundantly deferred calls.
        # https://go-critic.com/overview.html#unnecessarydefer
        - unnecessaryDefer
        # Detects slice expressions that can be simplified to sliced expression itself.
        # https://go-critic.com/overview.html#unslice
        - unslice
        # Detects value swapping code that are not using parallel assignment.
        # https://go-critic.com/overview.html#valswap
        - valSwap
        # Detects conditions that are unsafe due to not being exhaustive.
        # https://go-critic.com/overview.html#weakcond
        - weakCond
        # Ensures that //nolint comments include an explanation.
# https://go-critic.com/overview.html#whynolint
- whyNoLint
# Detects function calls that can be replaced with convenience wrappers.
# https://go-critic.com/overview.html#wrapperfunc
- wrapperFunc
# Detects Yoda style expressions and suggests to replace them.
# https://go-critic.com/overview.html#yodastyleexpr
- yodaStyleExpr

# Enable all checks.
# Default: false
enable-all: true
# Which checks should be disabled; can't be combined with 'enabled-checks'.
# Default: []
disabled-checks:
- appendAssign
- appendCombine
- argOrder
- assignOp
- badCall
- badCond
- badLock
- badRegexp
- badSorting
- badSyncOnceFunc
- boolExprSimplify
- builtinShadow
- builtinShadowDecl
- captLocal
- caseOrder
- codegenComment
- commentFormatting
- commentedOutCode
- commentedOutImport
- defaultCaseOrder
- deferInLoop
- deferUnlambda
- deprecatedComment
- docStub
- dupArg
- dupBranchBody
- dupCase
- dupImport
- dupSubExpr
- dynamicFmtString
- elseif
- emptyDecl
- emptyFallthrough
- emptyStringTest
- equalFold
- evalOrder
- exitAfterDefer
- exposedSyncMutex
- externalErrorReassign
- filepathJoin
- flagDeref
- flagName
- hexLiteral
- httpNoBody
- hugeParam
- ifElseChain
- importShadow
- indexAlloc
- initClause
- mapKey
- methodExprCall
- nestingReduce
- newDeref
- nilValReturn
- octalLiteral
- offBy1
- paramTypeCombine
- preferDecodeRune
- preferFilepathJoin
- preferFprint
- preferStringWriter
- preferWriteByte
- ptrToRefParam
- rangeAppendAll
- rangeExprCopy
- rangeValCopy
- redundantSprint
- regexpMust
- regexpPattern
- regexpSimplify
- returnAfterHttpError
- ruleguard
- singleCaseSwitch
- sliceClear
- sloppyLen
- sloppyReassign
- sloppyTypeAssert
- sortSlice
- sprintfQuotedString
- sqlQuery
- stringConcatSimplify
- stringXbytes
- stringsCompare
- switchTrue
- syncMapLoadAndDelete
- timeExprSimplify
- todoCommentWithoutDetail
- tooManyResultsChecker
- truncateCmp
- typeAssertChain
- typeDefFirst
- typeSwitchVar
- typeUnparen
- uncheckedInlineErr
- underef
- unlabelStmt
- unlambda
- unnamedResult
- unnecessaryBlock
- unnecessaryDefer
- unslice
- valSwap
- weakCond
- whyNoLint
- wrapperFunc
- yodaStyleExpr

# Enable multiple checks by tags in addition to default checks.
# Run GL_DEBUG=gocritic golangci-lint run --enable=gocritic to see all tags and checks.
# See https://github.com/go-critic/go-critic#usage -> section "Tags".
# Default: []
enabled-tags:
- diagnostic
- style
- performance
- experimental
- opinionated
disabled-tags:
- diagnostic
- style
- performance
- experimental
- opinionated

# Settings passed to gocritic.
# The settings key is the name of a supported gocritic checker.
# The list of supported checkers can be find in https://go-critic.com/overview.
settings:
# Must be valid enabled check name.
captLocal:
# Whether to restrict checker to params only.
# Default: true
paramsOnly: false
commentedOutCode:
# Min length of the comment that triggers a warning.
# Default: 15
minLength: 50
elseif:
# Whether to skip balanced if-else pairs.
# Default: true
skipBalanced: false
hugeParam:
# Size in bytes that makes the warning trigger.
# Default: 80
sizeThreshold: 70
ifElseChain:
# Min number of if-else blocks that makes the warning trigger.
# Default: 2
minThreshold: 4
nestingReduce:
# Min number of statements inside a branch to trigger a warning.
# Default: 5
bodyWidth: 4
rangeExprCopy:
# Size in bytes that makes the warning trigger.
# Default: 512
sizeThreshold: 516
# Whether to check test functions
# Default: true
skipTestFuncs: false
rangeValCopy:
# Size in bytes that makes the warning trigger.
# Default: 128
sizeThreshold: 32
# Whether to check test functions.
# Default: true
skipTestFuncs: false
ruleguard:
# Enable debug to identify which 'Where' condition was rejected.
# The value of the parameter is the name of a function in a ruleguard file.
#
# When a rule is evaluated:
# If:
#   The Match() clause is accepted; and
#   One of the conditions in the Where() clause is rejected,
# Then:
#   ruleguard prints the specific Where() condition that was rejected.
#
# The option is passed to the ruleguard 'debug-group' argument.
# Default: ""
debug: 'emptyDecl'
# Determines the behavior when an error occurs while parsing ruleguard files.
# If flag is not set, log error and skip rule files that contain an error.
# If flag is set, the value must be a comma-separated list of error conditions.
# - 'all':    fail on all errors.
# - 'import': ruleguard rule imports a package that cannot be found.
# - 'dsl':    gorule file does not comply with the ruleguard DSL.
# Default: ""
failOn: dsl,import
# Comma-separated list of file paths containing ruleguard rules.
# By default, if a path is relative, it is relative to the directory where the golangci-lint command is executed.
# The placeholder '${base-path}' is substituted with a path relative to the mode defined with run.relative-path-mode.
# Glob patterns such as 'rules-*.go' may be specified.
# Default: ""
rules: '${base-path}/ruleguard/rules-*.go,${base-path}/myrule1.go'
# Comma-separated list of enabled groups or skip empty to enable everything.
# Tags can be defined with # character prefix.
# Default: "<all>"
enable: "myGroupName,#myTagName"
# Comma-separated list of disabled groups or skip empty to enable everything.
# Tags can be defined with # character prefix.
# Default: ""
disable: "myGroupName,#myTagName"
tooManyResultsChecker:
# Maximum number of results.
# Default: 5
maxResults: 10
truncateCmp:
# Whether to skip int/uint/uintptr types.
# Default: true
skipArchDependent: false
underef:
# Whether to skip (*x).method() calls where x is a pointer receiver.
# Default: true
skipRecvDeref: false
unnamedResult:
# Whether to check exported functions.
# Default: false
checkExported: true

gocyclo:
# Minimal code complexity to report.
# Default: 30 (but we recommend 10-20)
min-complexity: 10

godot:
# Comments to be checked: declarations, toplevel, noinline or all.
# Default: declarations
scope: toplevel
# List of regexps for excluding particular comment lines from check.
# Default: []
exclude:
# Exclude todo and fixme comments.
- "^fixme:"
- "^todo:"
# Check that each sentence ends with a period.
# Default: true
period: false
# Check that each sentence starts with a capital letter.
# Default: false
capital: true

godox:
# Report any comments starting with keywords, this is useful for TODO or FIXME comments that
# might be left in the code accidentally and should be resolved before merging.
# Default: ["TODO", "BUG", "FIXME"]
keywords:
- NOTE
- OPTIMIZE # marks code that should be optimized before merging
- HACK # marks hack-around that should be removed before merging

goheader:
# Supports two types 'const and regexp.
# Values can be used recursively.
# Default: {}
values:
const:
# Define here const type values in format k:v.
# For example:
COMPANY: MY COMPANY
regexp:
# Define here regexp type values.
# for example:
AUTHOR: .*@mycompany\.com
# The template used for checking.
# Put here copyright header template for source code files
# Note: {{ YEAR }} is a builtin value that returns the year relative to the current machine time.
# Default: ""
template: |-
{{ AUTHOR }} {{ COMPANY }} {{ YEAR }}
SPDX-License-Identifier: Apache-2.0

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at:

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
# As alternative of directive 'template', you may put the path to file with the template source.
# Useful if you need to load the template from a specific file.
# By default, if a path is relative, it is relative to the directory where the golangci-lint command is executed.
# The placeholder '${base-path}' is substituted with a path relative to the mode defined with run.relative-path-mode.
# Default: ""
template-path: /path/to/my/template.tmpl

gomoddirectives:
# Allow local replace directives.
# Default: false
replace-local: true
# List of allowed replace directives.
# Default: []
replace-allow-list:
- launchpad.net/gocheck
# Allow to not explain why the version has been retracted in the retract directives.
# Default: false
retract-allow-no-explanation: true
# Forbid the use of the exclude directives.
# Default: false
exclude-forbidden: true
# Forbid the use of the toolchain directive.
# Default: false
toolchain-forbidden: true
# Defines a pattern to validate toolchain directive.
# Default: '' (no match)
toolchain-pattern: 'go1\.23\.\d+$'
# Forbid the use of the tool directives.
# Default: false
tool-forbidden: true
# Forbid the use of the godebug directive.
# Default: false
go-debug-forbidden: true
# Defines a pattern to validate go minimum version directive.
# Default: '' (no match)
go-version-pattern: '\d\.\d+(\.0)?'

gomodguard:
allowed:
# List of allowed modules.
# Default: []
modules:
- gopkg.in/yaml.v2
# List of allowed module domains.
# Default: []
domains:
- golang.org
blocked:
# List of blocked modules.
# Default: []
modules:
# Blocked module.
- github.com/uudashr/go-module:
# Recommended modules that should be used instead. (Optional)
recommendations:
- golang.org/x/mod
# Reason why the recommended module should be used. (Optional)
reason: "mod is the official go.mod parser library."
# List of blocked module version constraints.
# Default: []
versions:
# Blocked module with version constraint.
- github.com/mitchellh/go-homedir:
# Version constraint, see https://github.com/Masterminds/semver#basic-comparisons.
version: "< 1.1.0"
# Reason why the version constraint exists. (Optional)
reason: "testing if blocked version constraint works."
# Set to true to raise lint issues for packages that are loaded from a local path via replace directive.
# Default: false
local-replace-directives: false

gosec:
# To select a subset of rules to run.
# Available rules: https://github.com/securego/gosec#available-rules
# Default: [] - means include all rules
includes:
- G101 # Look for hard coded credentials
- G102 # Bind to all interfaces
- G103 # Audit the use of unsafe block
- G104 # Audit errors not checked
- G106 # Audit the use of ssh.InsecureIgnoreHostKey
- G107 # Url provided to HTTP request as taint input
- G108 # Profiling endpoint automatically exposed on /debug/pprof
- G109 # Potential Integer overflow made by strconv.Atoi result conversion to int16/32
- G110 # Potential DoS vulnerability via decompression bomb
- G111 # Potential directory traversal
- G112 # Potential slowloris attack
- G113 # Usage of Rat.SetString in math/big with an overflow (CVE-2022-23772)
- G114 # Use of net/http serve function that has no support for setting timeouts
- G115 # Potential integer overflow when converting between integer types
- G201 # SQL query construction using format string
- G202 # SQL query construction using string concatenation
- G203 # Use of unescaped data in HTML templates
- G204 # Audit use of command execution
- G301 # Poor file permissions used when creating a directory
- G302 # Poor file permissions used with chmod
- G303 # Creating tempfile using a predictable path
- G304 # File path provided as taint input
- G305 # File traversal when extracting zip/tar archive
- G306 # Poor file permissions used when writing to a new file
- G307 # Poor file permissions used when creating a file with os.Create
- G401 # Detect the usage of MD5 or SHA1
- G402 # Look for bad TLS connection settings
- G403 # Ensure minimum RSA key length of 2048 bits
- G404 # Insecure random number source (rand)
- G405 # Detect the usage of DES or RC4
- G406 # Detect the usage of MD4 or RIPEMD160
- G501 # Import blocklist: crypto/md5
- G502 # Import blocklist: crypto/des
- G503 # Import blocklist: crypto/rc4
- G504 # Import blocklist: net/http/cgi
- G505 # Import blocklist: crypto/sha1
- G506 # Import blocklist: golang.org/x/crypto/md4
- G507 # Import blocklist: golang.org/x/crypto/ripemd160
- G601 # Implicit memory aliasing of items from a range statement
- G602 # Slice access out of bounds

# To specify a set of rules to explicitly exclude.
# Available rules: https://github.com/securego/gosec#available-rules
# Default: []
excludes:
- G101 # Look for hard coded credentials
- G102 # Bind to all interfaces
- G103 # Audit the use of unsafe block
- G104 # Audit errors not checked
- G106 # Audit the use of ssh.InsecureIgnoreHostKey
- G107 # Url provided to HTTP request as taint input
- G108 # Profiling endpoint automatically exposed on /debug/pprof
- G109 # Potential Integer overflow made by strconv.Atoi result conversion to int16/32
- G110 # Potential DoS vulnerability via decompression bomb
- G111 # Potential directory traversal
- G112 # Potential slowloris attack
- G113 # Usage of Rat.SetString in math/big with an overflow (CVE-2022-23772)
- G114 # Use of net/http serve function that has no support for setting timeouts
- G115 # Potential integer overflow when converting between integer types
- G201 # SQL query construction using format string
- G202 # SQL query construction using string concatenation
- G203 # Use of unescaped data in HTML templates
- G204 # Audit use of command execution
- G301 # Poor file permissions used when creating a directory
- G302 # Poor file permissions used with chmod
- G303 # Creating tempfile using a predictable path
- G304 # File path provided as taint input
- G305 # File traversal when extracting zip/tar archive
- G306 # Poor file permissions used when writing to a new file
- G307 # Poor file permissions used when creating a file with os.Create
- G401 # Detect the usage of MD5 or SHA1
- G402 # Look for bad TLS connection settings
- G403 # Ensure minimum RSA key length of 2048 bits
- G404 # Insecure random number source (rand)
- G405 # Detect the usage of DES or RC4
- G406 # Detect the usage of MD4 or RIPEMD160
- G501 # Import blocklist: crypto/md5
- G502 # Import blocklist: crypto/des
- G503 # Import blocklist: crypto/rc4
- G504 # Import blocklist: net/http/cgi
- G505 # Import blocklist: crypto/sha1
- G506 # Import blocklist: golang.org/x/crypto/md4
- G507 # Import blocklist: golang.org/x/crypto/ripemd160
- G601 # Implicit memory aliasing of items from a range statement
- G602 # Slice access out of bounds

# Filter out the issues with a lower severity than the given value.
# Valid options are: low, medium, high.
# Default: low
severity: medium

# Filter out the issues with a lower confidence than the given value.
# Valid options are: low, medium, high.
# Default: low
confidence: medium

# Concurrency value.
# Default: the number of logical CPUs usable by the current process.
concurrency: 12

# To specify the configuration of rules.
config:
# Globals are applicable to all rules.
global:
# If true, ignore #nosec in comments (and an alternative as well).
# Default: false
nosec: true
# Add an alternative comment prefix to #nosec (both will work at the same time).
# Default: ""
"#nosec": "#my-custom-nosec"
# Define whether nosec issues are counted as finding or not.
# Default: false
show-ignored: true
# Audit mode enables addition checks that for normal code analysis might be too nosy.
# Default: false
audit: true
G101:
# Regexp pattern for variables and constants to find.
# Default: "(?i)passwd|pass|password|pwd|secret|token|pw|apiKey|bearer|cred"
pattern: "(?i)example"
# If true, complain about all cases (even with low entropy).
# Default: false
ignore_entropy: false
# Maximum allowed entropy of the string.
# Default: "80.0"
entropy_threshold: "80.0"
# Maximum allowed value of entropy/string length.
# Is taken into account if entropy >= entropy_threshold/2.
# Default: "3.0"
per_char_threshold: "3.0"
# Calculate entropy for first N chars of the string.
# Default: "16"
truncate: "32"
# Additional functions to ignore while checking unhandled errors.
# Following functions always ignored:
#   bytes.Buffer:
#     - Write
#     - WriteByte
#     - WriteRune
#     - WriteString
#   fmt:
#     - Print
#     - Printf
#     - Println
#     - Fprint
#     - Fprintf
#     - Fprintln
#   strings.Builder:
#     - Write
#     - WriteByte
#     - WriteRune
#     - WriteString
#   io.PipeWriter:
#     - CloseWithError
#   hash.Hash:
#     - Write
#   os:
#     - Unsetenv
# Default: {}
G104:
fmt:
- Fscanf
G111:
# Regexp pattern to find potential directory traversal.
# Default: "http\\.Dir\\(\"\\/\"\\)|http\\.Dir\\('\\/'\\)"
pattern: "custom\\.Dir\\(\\)"
# Maximum allowed permissions mode for os.Mkdir and os.MkdirAll
# Default: "0750"
G301: "0750"
# Maximum allowed permissions mode for os.OpenFile and os.Chmod
# Default: "0600"
G302: "0600"
# Maximum allowed permissions mode for os.WriteFile and ioutil.WriteFile
# Default: "0600"
G306: "0600"

gosmopolitan:
# Allow and ignore time.Local usages.
#
# Default: false
allow-time-local: true
# List of fully qualified names in the full/pkg/path.name form, to act as "i18n escape hatches".
# String literals inside call-like expressions to, or struct literals of those names,
# are exempt from the writing system check.
#
# Default: []
escape-hatches:
- 'github.com/nicksnyder/go-i18n/v2/i18n.Message'
- 'example.com/your/project/i18n/markers.Raw'
- 'example.com/your/project/i18n/markers.OK'
- 'example.com/your/project/i18n/markers.TODO'
- 'command-line-arguments.Simple'
# List of Unicode scripts to watch for any usage in string literals.
# https://pkg.go.dev/unicode#pkg-variables
#
# Default: ["Han"]
watch-for-scripts:
- Devanagari
- Han
- Hangul
- Hiragana
- Katakana

govet:
# Disable all analyzers.
# Default: false
disable-all: true
# Enable analyzers by name.
# (in addition to default:
#   appends, asmdecl, assign, atomic, bools, buildtag, cgocall, composites, copylocks, defers, directive, errorsas,
#   framepointer, httpresponse, ifaceassert, loopclosure, lostcancel, nilfunc, printf, shift, sigchanyzer, slog,
#   stdmethods, stringintconv, structtag, testinggoroutine, tests, timeformat, unmarshal, unreachable, unsafeptr,
#   unusedresult
# ).
# Run GL_DEBUG=govet golangci-lint run --enable=govet to see default, all available analyzers, and enabled analyzers.
# Default: []
enable:
# Check for missing values after append.
- appends
# Report mismatches between assembly files and Go declarations.
- asmdecl
# Check for useless assignments.
- assign
# Check for common mistakes using the sync/atomic package.
- atomic
# Check for non-64-bits-aligned arguments to sync/atomic functions.
- atomicalign
# Check for common mistakes involving boolean operators.
- bools
# Check //go:build and // +build directives.
- buildtag
# Detect some violations of the cgo pointer passing rules.
- cgocall
# Check for unkeyed composite literals.
- composites
# Check for locks erroneously passed by value.
- copylocks
# Check for calls of reflect.DeepEqual on error values.
- deepequalerrors
# Report common mistakes in defer statements.
- defers
# Check Go toolchain directives such as //go:debug.
- directive
# Report passing non-pointer or non-error values to errors.As.
- errorsas
# Find structs that would use less memory if their fields were sorted.
- fieldalignment
# Find calls to a particular function.
- findcall
# Report assembly that clobbers the frame pointer before saving it.
- framepointer
# Check for mistakes using HTTP responses.
- httpresponse
# Detect impossible interface-to-interface type assertions.
- ifaceassert
# Check references to loop variables from within nested functions.
- loopclosure
# Check cancel func returned by context.WithCancel is called.
- lostcancel
# Check for useless comparisons between functions and nil.
- nilfunc
# Check for redundant or impossible nil comparisons.
- nilness
# Check consistency of Printf format strings and arguments.
- printf
# Check for comparing reflect.Value values with == or reflect.DeepEqual.
- reflectvaluecompare
# Check for possible unintended shadowing of variables.
- shadow
# Check for shifts that equal or exceed the width of the integer.
- shift
# Check for unbuffered channel of os.Signal.
- sigchanyzer
# Check for invalid structured logging calls.
- slog
# Check the argument type of sort.Slice.
- sortslice
# Check signature of methods of well-known interfaces.
- stdmethods
# Report uses of too-new standard library symbols.
- stdversion
# Check for string(int) conversions.
- stringintconv
# Check that struct field tags conform to reflect.StructTag.Get.
- structtag
# Report calls to (*testing.T).Fatal from goroutines started by a test.
- testinggoroutine
# Check for common mistaken usages of tests and examples.
- tests
# Check for calls of (time.Time).Format or time.Parse with 2006-02-01.
- timeformat
# Report passing non-pointer or non-interface values to unmarshal.
- unmarshal
# Check for unreachable code.
- unreachable
# Check for invalid conversions of uintptr to unsafe.Pointer.
- unsafeptr
# Check for unused results of calls to some functions.
- unusedresult
# Checks for unused writes.
- unusedwrite
# Check for misuses of sync.WaitGroup.
- waitgroup

# Enable all analyzers.
# Default: false
enable-all: true
# Disable analyzers by name.
# (in addition to default
#   atomicalign, deepequalerrors, fieldalignment, findcall, nilness, reflectvaluecompare, shadow, sortslice,
#   timeformat, unusedwrite
# ).
# Run GL_DEBUG=govet golangci-lint run --enable=govet to see default, all available analyzers, and enabled analyzers.
# Default: []
disable:
- appends
- asmdecl
- assign
- atomic
- atomicalign
- bools
- buildtag
- cgocall
- composites
- copylocks
- deepequalerrors
- defers
- directive
- errorsas
- fieldalignment
- findcall
- framepointer
- httpresponse
- ifaceassert
- loopclosure
- lostcancel
- nilfunc
- nilness
- printf
- reflectvaluecompare
- shadow
- shift
- sigchanyzer
- slog
- sortslice
- stdmethods
- stdversion
- stringintconv
- structtag
- testinggoroutine
- tests
- timeformat
- unmarshal
- unreachable
- unsafeptr
- unusedresult
- unusedwrite
- waitgroup

# Settings per analyzer.
settings:
# Analyzer name, run go tool vet help to see all analyzers.
printf:
# Comma-separated list of print function names to check (in addition to default, see go tool vet help printf).
# Default: []
funcs:
- (github.com/golangci/golangci-lint/v2/pkg/logutils.Log).Infof
- (github.com/golangci/golangci-lint/v2/pkg/logutils.Log).Warnf
- (github.com/golangci/golangci-lint/v2/pkg/logutils.Log).Errorf
- (github.com/golangci/golangci-lint/v2/pkg/logutils.Log).Fatalf
shadow:
# Whether to be strict about shadowing; can be noisy.
# Default: false
strict: true
unusedresult:
# Comma-separated list of functions whose results must be used
# (in addition to default:
#   context.WithCancel, context.WithDeadline, context.WithTimeout, context.WithValue, errors.New, fmt.Errorf,
#   fmt.Sprint, fmt.Sprintf, sort.Reverse
# ).
# Default: []
funcs:
- pkg.MyFunc
# Comma-separated list of names of methods of type func() string whose results must be used
# (in addition to default Error,String)
# Default: []
stringmethods:
- MyMethod

grouper:
# Require the use of a single global 'const' declaration only.
# Default: false
const-require-single-const: true
# Require the use of grouped global 'const' declarations.
# Default: false
const-require-grouping: true

# Require the use of a single 'import' declaration only.
# Default: false
import-require-single-import: true
# Require the use of grouped 'import' declarations.
# Default: false
import-require-grouping: true

# Require the use of a single global 'type' declaration only.
# Default: false
type-require-single-type: true
# Require the use of grouped global 'type' declarations.
# Default: false
type-require-grouping: true

# Require the use of a single global 'var' declaration only.
# Default: false
var-require-single-var: true
# Require the use of grouped global 'var' declarations.
# Default: false
var-require-grouping: true

iface:
# List of analyzers.
# Default: ["identical"]
enable:
- identical # Identifies interfaces in the same package that have identical method sets.
- unused # Identifies interfaces that are not used anywhere in the same package where the interface is defined.
- opaque # Identifies functions that return interfaces, but the actual returned value is always a single concrete implementation.
settings:
unused:
# List of packages path to exclude from the check.
# Default: []
exclude:
- github.com/example/log

importas:
# Do not allow unaliased imports of aliased packages.
# Default: false
no-unaliased: true
# Do not allow non-required aliases.
# Default: false
no-extra-aliases: true
# List of aliases
# Default: []
alias:
# Using servingv1 alias for knative.dev/serving/pkg/apis/serving/v1 package.
- pkg: knative.dev/serving/pkg/apis/serving/v1
alias: servingv1
# Using autoscalingv1alpha1 alias for knative.dev/serving/pkg/apis/autoscaling/v1alpha1 package.
- pkg: knative.dev/serving/pkg/apis/autoscaling/v1alpha1
alias: autoscalingv1alpha1
# You can specify the package path by regular expression,
# and alias by regular expression expansion syntax like below.
# see https://github.com/julz/importas#use-regular-expression for details
- pkg: knative.dev/serving/pkg/apis/(\w+)/(v[\w\d]+)
alias: $1$2
# An explicit empty alias can be used to ensure no aliases are used for a package.
# This can be useful if no-extra-aliases: true doesn't fit your need.
# Multiple packages can use an empty alias.
- pkg: errors
alias: ""

inamedparam:
# Skips check for interface methods with only a single parameter.
# Default: false
skip-single-param: true

interfacebloat:
# The maximum number of methods allowed for an interface.
# Default: 10
max: 5

ireturn:
# List of interfaces to allow.
# Lists of the keywords and regular expressions matched to interface or package names can be used.
# allow and reject settings cannot be used at the same time.
#
# Keywords:
# - empty for interface{}
# - error for errors
# - stdlib for standard library
# - anon for anonymous interfaces
# - generic for generic interfaces added in go 1.18
#
# Default: [anon, error, empty, stdlib]
allow:
- anon
# You can specify idiomatic endings for interface
- (or|er)$

# List of interfaces to reject.
# Lists of the keywords and regular expressions matched to interface or package names can be used.
# allow and reject settings cannot be used at the same time.
#
# Keywords:
# - empty for interface{}
# - error for errors
# - stdlib for standard library
# - anon for anonymous interfaces
# - generic for generic interfaces added in go 1.18
#
# Default: []
reject:
- github.com\/user\/package\/v4\.Type

lll:
# Max line length, lines longer will be reported.
# '\t' is counted as 1 character by default, and can be changed with the tab-width option.
# Default: 120.
line-length: 120
# Tab width in spaces.
# Default: 1
tab-width: 1

loggercheck:
# Allow check for the github.com/go-kit/log library.
# Default: true
kitlog: false
# Allow check for the k8s.io/klog/v2 library.
# Default: true
klog: false
# Allow check for the github.com/go-logr/logr library.
# Default: true
logr: false
# Allow check for the log/slog library.
# Default: true
slog: false
# Allow check for the "sugar logger" from go.uber.org/zap library.
# Default: true
zap: false
# Require all logging keys to be inlined constant strings.
# Default: false
require-string-key: true
# Require printf-like format specifier (%s, %d for example) not present.
# Default: false
no-printf-like: true
# List of custom rules to check against, where each rule is a single logger pattern, useful for wrapped loggers.
# For example: https://github.com/timonwong/loggercheck/blob/7395ab86595781e33f7afba27ad7b55e6956ebcd/testdata/custom-rules.txt
# Default: empty
rules:
- k8s.io/klog/v2.InfoS   # package level exported functions
- (github.com/go-logr/logr.Logger).Error  # "Methods"
- (*go.uber.org/zap.SugaredLogger).With  # Also "Methods", but with a pointer receiver

maintidx:
# Show functions with maintainability index lower than N.
# A high index indicates better maintainability (it's kind of the opposite of complexity).
# Default: 20
under: 100

makezero:
# Allow only slices initialized with a length of zero.
# Default: false
always: true

misspell:
# Correct spellings using locale preferences for US or UK.
# Setting locale to US will correct the British spelling of 'colour' to 'color'.
# Default is to use a neutral variety of English.
locale: US
# Typos to ignore.
# Should be in lower case.
# Default: []
ignore-rules:
- someword
# Extra word corrections.
# typo and correction should only contain letters.
# The words are case-insensitive.
# Default: []
extra-words:
- typo: "iff"
correction: "if"
- typo: "cancelation"
correction: "cancellation"
# Mode of the analysis:
# - default: checks all the file content.
# - restricted: checks only comments.
# Default: ""
mode: restricted

mnd:
# List of enabled checks, see https://github.com/tommy-muehle/go-mnd/#checks for description.
# Default: ["argument", "case", "condition", "operation", "return", "assign"]
checks:
- argument
- case
- condition
- operation
- return
- assign
# List of numbers to exclude from analysis.
# The numbers should be written as string.
# Values always ignored: "1", "1.0", "0" and "0.0"
# Default: []
ignored-numbers:
- '0666'
- '0755'
- '42'
# List of file patterns to exclude from analysis.
# Values always ignored: .+_test.go
# Default: []
ignored-files:
- 'magic1_.+\.go$'
# List of function patterns to exclude from analysis.
# Following functions are always ignored: time.Date,
# strconv.FormatInt, strconv.FormatUint, strconv.FormatFloat,
# strconv.ParseInt, strconv.ParseUint, strconv.ParseFloat.
# Default: []
ignored-functions:
- '^math\.'
- '^http\.StatusText$'

musttag:
# A set of custom functions to check in addition to the builtin ones.
# Default: json, xml, gopkg.in/yaml.v3, BurntSushi/toml, mitchellh/mapstructure, jmoiron/sqlx
functions:
# The full name of the function, including the package.
- name: github.com/hashicorp/hcl/v2/hclsimple.DecodeFile
# The struct tag whose presence should be ensured.
tag: hcl
# The position of the argument to check.
arg-pos: 2

nakedret:
# Make an issue if func has more lines of code than this setting, and it has naked returns.
# Default: 30
max-func-lines: 31

nestif:
# Minimal complexity of if statements to report.
# Default: 5
min-complexity: 4

nilnil:
# To check functions with only two return values (return nil, nil).
# If disabled then returns like return nil, nil, ..., nil are supported.
# Default: true
only-two: false
# In addition, detect opposite situation (simultaneous return of non-nil error and valid value).
# E.g, return clone, fh.indexer.Update(clone) will be considered as invalid.
# Default: false
detect-opposite: true
# List of return types to check.
# Default: ["chan", "func", "iface", "map", "ptr", "uintptr", "unsafeptr"]
checked-types:
- chan
- func
- iface
- map
- ptr
- uintptr
- unsafeptr

nlreturn:
# Size of the block (including return statement that is still "OK")
# so no return split required.
# Default: 1
block-size: 2

nolintlint:
# Disable to ensure that all nolint directives actually have an effect.
# Default: false
allow-unused: true
# Exclude following linters from requiring an explanation.
# Default: []
allow-no-explanation: [ ]
# Enable to require an explanation of nonzero length after each nolint directive.
# Default: false
require-explanation: true
# Enable to require nolint directives to mention the specific linter being suppressed.
# Default: false
require-specific: true

nonamedreturns:
# Report named error if it is assigned inside defer.
# Default: false
report-error-in-defer: true

paralleltest:
# Ignore missing calls to t.Parallel() and only report incorrect uses of it.
# Default: false
ignore-missing: true
# Ignore missing calls to t.Parallel() in subtests. Top-level tests are
# still required to have t.Parallel, but subtests are allowed to skip it.
# Default: false
ignore-missing-subtests: true

perfsprint:
# Enable/disable optimization of integer formatting.
# Default: true
integer-format: false
# Optimizes even if it requires an int or uint type cast.
# Default: true
int-conversion: false
# Enable/disable optimization of error formatting.
# Default: true
error-format: false
# Optimizes into err.Error() even if it is only equivalent for non-nil errors.
# Default: false
err-error: true
# Optimizes fmt.Errorf.
# Default: true
errorf: false
# Enable/disable optimization of string formatting.
# Default: true
string-format: false
# Optimizes fmt.Sprintf with only one argument.
# Default: true
sprintf1: false
# Optimizes into strings concatenation.
# Default: true
strconcat: false
# Enable/disable optimization of bool formatting.
# Default: true
bool-format: false
# Enable/disable optimization of hex formatting.
# Default: true
hex-format: false

prealloc:
# IMPORTANT: we don't recommend using this linter before doing performance profiling.
# For most programs usage of prealloc will be a premature optimization.

# Report pre-allocation suggestions only on simple loops that have no returns/breaks/continues/gotos in them.
# Default: true
simple: false
# Report pre-allocation suggestions on range loops.
# Default: true
range-loops: false
# Report pre-allocation suggestions on for loops.
# Default: false
for-loops: true

predeclared:
# List of predeclared identifiers to not report on.
# Default: []
ignore:
- new
- int
# Include method names and field names in checks.
# Default: false
qualified-name: true

promlinter:
# Promlinter cannot infer all metrics name in static analysis.
# Enable strict mode will also include the errors caused by failing to parse the args.
# Default: false
strict: true
# Please refer to https://github.com/yeya24/promlinter#usage for detailed usage.
# Default: []
disabled-linters:
# Help detects issues related to the help text for a metric.
- Help
# MetricUnits detects issues with metric unit names.
- MetricUnits
# Counter detects issues specific to counters, as well as patterns that should only be used with counters.
- Counter
# HistogramSummaryReserved detects when other types of metrics use names or labels reserved for use by histograms and/or summaries.
- HistogramSummaryReserved
# MetricTypeInName detects when metric types are included in the metric name.
- MetricTypeInName
# ReservedChars detects colons in metric names.
- ReservedChars
# CamelCase detects metric names and label names written in camelCase.
- CamelCase
# UnitAbbreviations detects abbreviated units in the metric name.
- UnitAbbreviations

protogetter:
# Skip files generated by specified generators from the checking.
# Checks only the file's initial comment, which must follow the format: "// Code generated by <generator-name>".
# Files generated by protoc-gen-go, protoc-gen-go-grpc, and protoc-gen-grpc-gateway are always excluded automatically.
# Default: []
skip-generated-by: ["protoc-gen-go-my-own-generator"]
# Skip files matching the specified glob pattern from the checking.
# Default: []
skip-files:
- "*.pb.go"
- "*/vendor/*"
- "/full/path/to/file.go"
# Skip any generated files from the checking.
# Default: false
skip-any-generated: true
# Skip first argument of append function.
# Default: false
replace-first-arg-in-append: true

reassign:
# Patterns for global variable names that are checked for reassignment.
# See https://github.com/curioswitch/go-reassign#usage
# Default: ["EOF", "Err.*"]
patterns:
- ".*"

recvcheck:
# Disables the built-in method exclusions:
# - MarshalText
# - MarshalJSON
# - MarshalYAML
# - MarshalXML
# - MarshalBinary
# - GobEncode
# Default: false
disable-builtin: true
# User-defined method exclusions.
# The format is struct_name.method_name (ex: Foo.MethodName).
# A wildcard * can use as a struct name (ex: *.MethodName).
# Default: []
exclusions:
- "*.Value"

revive:
# Maximum number of open files at the same time.
# See https://github.com/mgechev/revive#command-line-flags
# Defaults to unlimited.
max-open-files: 2048

# Sets the default severity.
# See https://github.com/mgechev/revive#configuration
# Default: warning
severity: error

# Enable all available rules.
# Default: false
enable-all-rules: true

# Enable validation of comment directives.
# See https://github.com/mgechev/revive#comment-directives
directives:
- name: specify-disable-reason
severity: error

# Sets the default failure confidence.
# This means that linting errors with less than 0.8 confidence will be ignored.
# Default: 0.8
confidence: 0.1
# Run GL_DEBUG=revive golangci-lint run --enable-only=revive to see default, all available rules, and enabled rules.
rules:
# https://github.com/mgechev/revive/blob/HEAD/RULES_DESCRIPTIONS.md#add-constant
- name: add-constant
severity: warning
disabled: false
exclude: [""]
arguments:
- maxLitCount: "3"
allowStrs: '""'
allowInts: "0,1,2"
allowFloats: "0.0,0.,1.0,1.,2.0,2."
# https://github.com/mgechev/revive/blob/HEAD/RULES_DESCRIPTIONS.md#argument-limit
- name: argument-limit
severity: warning
disabled: false
exclude: [""]
arguments: [ 4 ]
# https://github.com/mgechev/revive/blob/HEAD/RULES_DESCRIPTIONS.md#atomic
- name: atomic
severity: warning
disabled: false
exclude: [""]
# https://github.com/mgechev/revive/blob/HEAD/RULES_DESCRIPTIONS.md#banned-characters
- name: banned-characters
severity: warning
disabled: false
exclude: [""]
arguments: [ "Ω","Σ","σ", "7" ]
# https://github.com/mgechev/revive/blob/HEAD/RULES_DESCRIPTIONS.md#bare-return
- name: bare-return
severity: warning
disabled: false
exclude: [""]
# https://github.com/mgechev/revive/blob/HEAD/RULES_DESCRIPTIONS.md#blank-imports
- name: blank-imports
severity: warning
disabled: false
exclude: [""]
# https://github.com/mgechev/revive/blob/HEAD/RULES_DESCRIPTIONS.md#bool-literal-in-expr
- name: bool-literal-in-expr
severity: warning
disabled: false
exclude: [""]
# https://github.com/mgechev/revive/blob/HEAD/RULES_DESCRIPTIONS.md#call-to-gc
- name: call-to-gc
severity: warning
disabled: false
exclude: [""]
# https://github.com/mgechev/revive/blob/HEAD/RULES_DESCRIPTIONS.md#cognitive-complexity
- name: cognitive-complexity
severity: warning
disabled: false
exclude: [""]
arguments: [ 7 ]
# https://github.com/mgechev/revive/blob/HEAD/RULES_DESCRIPTIONS.md#comment-spacings
- name: comment-spacings
severity: warning
disabled: false
exclude: [""]
arguments:
- mypragma
- otherpragma
# https://github.com/mgechev/revive/blob/HEAD/RULES_DESCRIPTIONS.md#comments-density
- name: comments-density
severity: warning
disabled: false
exclude: [""]
arguments: [ 15 ]
# https://github.com/mgechev/revive/blob/HEAD/RULES_DESCRIPTIONS.md#confusing-naming
- name: confusing-naming
severity: warning
disabled: false
exclude: [""]
# https://github.com/mgechev/revive/blob/HEAD/RULES_DESCRIPTIONS.md#confusing-results
- name: confusing-results
severity: warning
disabled: false
exclude: [""]
# https://github.com/mgechev/revive/blob/HEAD/RULES_DESCRIPTIONS.md#constant-logical-expr
- name: constant-logical-expr
severity: warning
disabled: false
exclude: [""]
# https://github.com/mgechev/revive/blob/HEAD/RULES_DESCRIPTIONS.md#context-as-argument
- name: context-as-argument
severity: warning
disabled: false
exclude: [""]
arguments:
- allowTypesBefore: "*testing.T,*github.com/user/repo/testing.Harness"
# https://github.com/mgechev/revive/blob/HEAD/RULES_DESCRIPTIONS.md#context-keys-type
- name: context-keys-type
	severity: warning
disabled: false
exclude: [""]
# https://github.com/mgechev/revive/blob/HEAD/RULES_DESCRIPTIONS.md#cyclomatic
- name: cyclomatic
severity: warning
disabled: false
exclude: [""]
arguments: [ 3 ]
# https://github.com/mgechev/revive/blob/HEAD/RULES_DESCRIPTIONS.md#datarace
- name: datarace
severity: warning
disabled: false
exclude: [""]
# https://github.com/mgechev/revive/blob/HEAD/RULES_DESCRIPTIONS.md#deep-exit
- name: deep-exit
severity: warning
disabled: false
exclude: [""]
# https://github.com/mgechev/revive/blob/HEAD/RULES_DESCRIPTIONS.md#defer
- name: defer
severity: warning
disabled: false
exclude: [""]
arguments:
- "call-chain"
- "loop"
- "method-call"
- "recover"
- "immediate-recover"
- "return"
# https://github.com/mgechev/revive/blob/HEAD/RULES_DESCRIPTIONS.md#dot-imports
- name: dot-imports
severity: warning
disabled: false
exclude: [""]
arguments:
- allowedPackages: ["github.com/onsi/ginkgo/v2", "github.com/onsi/gomega"]
# https://github.com/mgechev/revive/blob/HEAD/RULES_DESCRIPTIONS.md#duplicated-imports
- name: duplicated-imports
severity: warning
disabled: false
exclude: [""]
# https://github.com/mgechev/revive/blob/HEAD/RULES_DESCRIPTIONS.md#early-return
- name: early-return
severity: warning
disabled: false
exclude: [""]
arguments:
- "preserveScope"
- "allowJump"
# https://github.com/mgechev/revive/blob/HEAD/RULES_DESCRIPTIONS.md#empty-block
- name: empty-block
severity: warning
disabled: false
exclude: [""]
# https://github.com/mgechev/revive/blob/HEAD/RULES_DESCRIPTIONS.md#empty-lines
- name: empty-lines
severity: warning
disabled: false
exclude: [""]
# https://github.com/mgechev/revive/blob/HEAD/RULES_DESCRIPTIONS.md#enforce-map-style
- name: enforce-map-style
severity: warning
disabled: false
exclude: [""]
arguments:
- "make"
# https://github.com/mgechev/revive/blob/HEAD/RULES_DESCRIPTIONS.md#enforce-repeated-arg-type-style
- name: enforce-repeated-arg-type-style
severity: warning
disabled: false
exclude: [""]
arguments:
- "short"
# Or this parameter:
- funcArgStyle: "full"
funcRetValStyle: "short"
# https://github.com/mgechev/revive/blob/HEAD/RULES_DESCRIPTIONS.md#enforce-slice-style
- name: enforce-slice-style
severity: warning
disabled: false
exclude: [""]
arguments:
- "make"
# https://github.com/mgechev/revive/blob/HEAD/RULES_DESCRIPTIONS.md#error-naming
- name: error-naming
severity: warning
disabled: false
exclude: [""]
# https://github.com/mgechev/revive/blob/HEAD/RULES_DESCRIPTIONS.md#error-return
- name: error-return
severity: warning
disabled: false
exclude: [""]
# https://github.com/mgechev/revive/blob/HEAD/RULES_DESCRIPTIONS.md#error-strings
- name: error-strings
severity: warning
disabled: false
exclude: [""]
arguments:
- "xerrors.New"
# https://github.com/mgechev/revive/blob/HEAD/RULES_DESCRIPTIONS.md#errorf
- name: errorf
severity: warning
disabled: false
exclude: [""]
# https://github.com/mgechev/revive/blob/HEAD/RULES_DESCRIPTIONS.md#exported
- name: exported
severity: warning
disabled: false
exclude: [""]
arguments:
- "checkPrivateReceivers"
- "disableStutteringCheck"
- "sayRepetitiveInsteadOfStutters"
- "checkPublicInterface"
- "disableChecksOnConstants"
- "disableChecksOnFunctions"
- "disableChecksOnMethods"
- "disableChecksOnTypes"
- "disableChecksOnVariables"
# https://github.com/mgechev/revive/blob/HEAD/RULES_DESCRIPTIONS.md#file-header
- name: file-header
severity: warning
disabled: false
exclude: [""]
arguments:
- This is the text that must appear at the top of source files.
# https://github.com/mgechev/revive/blob/HEAD/RULES_DESCRIPTIONS.md#file-length-limit
- name: file-length-limit
severity: warning
disabled: false
exclude: [""]
arguments:
- max: 100
skipComments: true
skipBlankLines: true
# https://github.com/mgechev/revive/blob/HEAD/RULES_DESCRIPTIONS.md#filename-format
- name: filename-format
severity: warning
disabled: false
exclude: [""]
arguments:
- "^[_a-z][_a-z0-9]*\\.go$"
# https://github.com/mgechev/revive/blob/HEAD/RULES_DESCRIPTIONS.md#flag-parameter
- name: flag-parameter
severity: warning
disabled: false
exclude: [""]
# https://github.com/mgechev/revive/blob/HEAD/RULES_DESCRIPTIONS.md#function-length
- name: function-length
severity: warning
disabled: false
exclude: [""]
arguments: [ 10, 0 ]
# https://github.com/mgechev/revive/blob/HEAD/RULES_DESCRIPTIONS.md#function-result-limit
- name: function-result-limit
severity: warning
disabled: false
exclude: [""]
arguments: [ 3 ]
# https://github.com/mgechev/revive/blob/HEAD/RULES_DESCRIPTIONS.md#get-return
- name: get-return
severity: warning
disabled: false
exclude: [""]
# https://github.com/mgechev/revive/blob/HEAD/RULES_DESCRIPTIONS.md#identical-branches
- name: identical-branches
severity: warning
disabled: false
exclude: [""]
# https://github.com/mgechev/revive/blob/HEAD/RULES_DESCRIPTIONS.md#if-return
- name: if-return
severity: warning
disabled: false
exclude: [""]
# https://github.com/mgechev/revive/blob/HEAD/RULES_DESCRIPTIONS.md#import-alias-naming
- name: import-alias-naming
severity: warning
disabled: false
exclude: [""]
arguments:
- "^[a-z][a-z0-9]{0,}$"
# Or this parameter:
- allowRegex: "^[a-z][a-z0-9]{0,}$"
denyRegex: '^v\d+$'
# https://github.com/mgechev/revive/blob/HEAD/RULES_DESCRIPTIONS.md#import-shadowing
- name: import-shadowing
severity: warning
disabled: false
exclude: [""]
# https://github.com/mgechev/revive/blob/HEAD/RULES_DESCRIPTIONS.md#imports-blocklist
- name: imports-blocklist
severity: warning
disabled: false
exclude: [""]
arguments:
- "crypto/md5"
- "crypto/sha1"
# https://github.com/mgechev/revive/blob/HEAD/RULES_DESCRIPTIONS.md#increment-decrement
- name: increment-decrement
severity: warning
disabled: false
exclude: [""]
# https://github.com/mgechev/revive/blob/HEAD/RULES_DESCRIPTIONS.md#indent-error-flow
- name: indent-error-flow
severity: warning
disabled: false
exclude: [""]
arguments:
- "preserveScope"
# https://github.com/mgechev/revive/blob/HEAD/RULES_DESCRIPTIONS.md#line-length-limit
- name: line-length-limit
severity: warning
disabled: false
exclude: [""]
arguments: [ 80 ]
# https://github.com/mgechev/revive/blob/HEAD/RULES_DESCRIPTIONS.md#max-control-nesting
- name: max-control-nesting
severity: warning
disabled: false
exclude: [""]
arguments: [ 3 ]
# https://github.com/mgechev/revive/blob/HEAD/RULES_DESCRIPTIONS.md#max-public-structs
- name: max-public-structs
severity: warning
disabled: false
exclude: [""]
arguments: [ 3 ]
# https://github.com/mgechev/revive/blob/HEAD/RULES_DESCRIPTIONS.md#modifies-parameter
- name: modifies-parameter
severity: warning
disabled: false
exclude: [""]
# https://github.com/mgechev/revive/blob/HEAD/RULES_DESCRIPTIONS.md#modifies-value-receiver
- name: modifies-value-receiver
severity: warning
disabled: false
exclude: [""]
# https://github.com/mgechev/revive/blob/HEAD/RULES_DESCRIPTIONS.md#nested-structs
- name: nested-structs
severity: warning
disabled: false
exclude: [""]
# https://github.com/mgechev/revive/blob/HEAD/RULES_DESCRIPTIONS.md#optimize-operands-order
- name: optimize-operands-order
severity: warning
disabled: false
exclude: [""]
# https://github.com/mgechev/revive/blob/HEAD/RULES_DESCRIPTIONS.md#package-comments
- name: package-comments
severity: warning
disabled: false
exclude: [""]
# https://github.com/mgechev/revive/blob/HEAD/RULES_DESCRIPTIONS.md#range
- name: range
severity: warning
disabled: false
exclude: [""]
# https://github.com/mgechev/revive/blob/HEAD/RULES_DESCRIPTIONS.md#range-val-address
- name: range-val-address
severity: warning
disabled: false
exclude: [""]
# https://github.com/mgechev/revive/blob/HEAD/RULES_DESCRIPTIONS.md#range-val-in-closure
- name: range-val-in-closure
severity: warning
disabled: false
exclude: [""]
# https://github.com/mgechev/revive/blob/HEAD/RULES_DESCRIPTIONS.md#receiver-naming
- name: receiver-naming
severity: warning
disabled: false
exclude: [""]
arguments:
- maxLength: 2
# https://github.com/mgechev/revive/blob/HEAD/RULES_DESCRIPTIONS.md#redefines-builtin-id
- name: redefines-builtin-id
severity: warning
disabled: false
exclude: [""]
# https://github.com/mgechev/revive/blob/HEAD/RULES_DESCRIPTIONS.md#redundant-build-tag
- name: redundant-build-tag
severity: warning
disabled: false
exclude: [""]
# https://github.com/mgechev/revive/blob/HEAD/RULES_DESCRIPTIONS.md#redundant-import-alias
- name: redundant-import-alias
severity: warning
disabled: false
exclude: [""]
# https://github.com/mgechev/revive/blob/HEAD/RULES_DESCRIPTIONS.md#string-format
- name: string-format
severity: warning
disabled: false
exclude: [""]
arguments:
- - 'core.WriteError[1].Message'
- '/^([^A-Z]|$)/'
- must not start with a capital letter
- - 'fmt.Errorf[0]'
- '/(^|[^\.!?])$/'
- must not end in punctuation
- - panic
- '/^[^\n]*$/'
- must not contain line breaks
# https://github.com/mgechev/revive/blob/HEAD/RULES_DESCRIPTIONS.md#string-of-int
- name: string-of-int
severity: warning
disabled: false
exclude: [""]
# https://github.com/mgechev/revive/blob/HEAD/RULES_DESCRIPTIONS.md#struct-tag
- name: struct-tag
severity: warning
disabled: false
exclude: [""]
arguments:
- "json,inline"
- "bson,outline,gnu"
# https://github.com/mgechev/revive/blob/HEAD/RULES_DESCRIPTIONS.md#superfluous-else
- name: superfluous-else
severity: warning
disabled: false
exclude: [""]
arguments:
- "preserveScope"
# https://github.com/mgechev/revive/blob/HEAD/RULES_DESCRIPTIONS.md#time-equal
- name: time-equal
severity: warning
disabled: false
exclude: [""]
# https://github.com/mgechev/revive/blob/HEAD/RULES_DESCRIPTIONS.md#time-naming
- name: time-naming
severity: warning
disabled: false
exclude: [""]
# https://github.com/mgechev/revive/blob/HEAD/RULES_DESCRIPTIONS.md#unchecked-type-assertion
- name: unchecked-type-assertion
severity: warning
disabled: false
exclude: [""]
arguments:
- acceptIgnoredAssertionResult: true
# https://github.com/mgechev/revive/blob/HEAD/RULES_DESCRIPTIONS.md#unconditional-recursion
- name: unconditional-recursion
severity: warning
disabled: false
exclude: [""]
# https://github.com/mgechev/revive/blob/HEAD/RULES_DESCRIPTIONS.md#unexported-naming
- name: unexported-naming
severity: warning
disabled: false
exclude: [""]
# https://github.com/mgechev/revive/blob/HEAD/RULES_DESCRIPTIONS.md#unexported-return
- name: unexported-return
severity: warning
disabled: false
exclude: [""]
# https://github.com/mgechev/revive/blob/HEAD/RULES_DESCRIPTIONS.md#unhandled-error
- name: unhandled-error
severity: warning
disabled: false
exclude: [""]
arguments:
- "fmt.Printf"
- "myFunction"
# https://github.com/mgechev/revive/blob/HEAD/RULES_DESCRIPTIONS.md#unnecessary-stmt
- name: unnecessary-stmt
severity: warning
disabled: false
exclude: [""]
# https://github.com/mgechev/revive/blob/HEAD/RULES_DESCRIPTIONS.md#unreachable-code
- name: unreachable-code
severity: warning
disabled: false
exclude: [""]
# https://github.com/mgechev/revive/blob/HEAD/RULES_DESCRIPTIONS.md#unused-parameter
- name: unused-parameter
severity: warning
disabled: false
exclude: [""]
arguments:
- allowRegex: "^_"
# https://github.com/mgechev/revive/blob/HEAD/RULES_DESCRIPTIONS.md#unused-receiver
- name: unused-receiver
severity: warning
disabled: false
exclude: [""]
arguments:
- allowRegex: "^_"
# https://github.com/mgechev/revive/blob/HEAD/RULES_DESCRIPTIONS.md#use-any
- name: use-any
severity: warning
disabled: false
exclude: [""]
# https://github.com/mgechev/revive/blob/HEAD/RULES_DESCRIPTIONS.md#use-errors-new
- name: use-errors-new
severity: warning
disabled: false
exclude: [""]
# https://github.com/mgechev/revive/blob/HEAD/RULES_DESCRIPTIONS.md#useless-break
- name: useless-break
severity: warning
disabled: false
exclude: [""]
# https://github.com/mgechev/revive/blob/HEAD/RULES_DESCRIPTIONS.md#var-declaration
- name: var-declaration
severity: warning
disabled: false
exclude: [""]
# https://github.com/mgechev/revive/blob/HEAD/RULES_DESCRIPTIONS.md#var-naming
- name: var-naming
severity: warning
disabled: false
exclude: [""]
arguments:
- [ "ID" ] # AllowList
- [ "VM" ] # DenyList
- - upperCaseConst: true # Extra parameter (upperCaseConst|skipPackageNameChecks)
# https://github.com/mgechev/revive/blob/HEAD/RULES_DESCRIPTIONS.md#waitgroup-by-value
- name: waitgroup-by-value
severity: warning
disabled: false
exclude: [""]

rowserrcheck:
# database/sql is always checked
# Default: []
packages:
- github.com/jmoiron/sqlx

sloglint:
# Enforce not mixing key-value pairs and attributes.
# https://github.com/go-simpler/sloglint?tab=readme-ov-file#no-mixed-arguments
# Default: true
no-mixed-args: false
# Enforce using key-value pairs only (overrides no-mixed-args, incompatible with attr-only).
# https://github.com/go-simpler/sloglint?tab=readme-ov-file#key-value-pairs-only
# Default: false
kv-only: true
# Enforce using attributes only (overrides no-mixed-args, incompatible with kv-only).
# https://github.com/go-simpler/sloglint?tab=readme-ov-file#attributes-only
# Default: false
attr-only: true
# Enforce not using global loggers.
# Values:
# - "": disabled
# - "all": report all global loggers
# - "default": report only the default slog logger
# https://github.com/go-simpler/sloglint?tab=readme-ov-file#no-global
# Default: ""
no-global: "all"
# Enforce using methods that accept a context.
# Values:
# - "": disabled
# - "all": report all contextless calls
# - "scope": report only if a context exists in the scope of the outermost function
# https://github.com/go-simpler/sloglint?tab=readme-ov-file#context-only
# Default: ""
context: "all"
# Enforce using static values for log messages.
# https://github.com/go-simpler/sloglint?tab=readme-ov-file#static-messages
# Default: false
static-msg: true
# Enforce using constants instead of raw keys.
# https://github.com/go-simpler/sloglint?tab=readme-ov-file#no-raw-keys
# Default: false
no-raw-keys: true
# Enforce a single key naming convention.
# Values: snake, kebab, camel, pascal
# https://github.com/go-simpler/sloglint?tab=readme-ov-file#key-naming-convention
# Default: ""
key-naming-case: snake
# Enforce not using specific keys.
# https://github.com/go-simpler/sloglint?tab=readme-ov-file#forbidden-keys
# Default: []
forbidden-keys:
- time
- level
- msg
- source
- foo
# Enforce putting arguments on separate lines.
# https://github.com/go-simpler/sloglint?tab=readme-ov-file#arguments-on-separate-lines
# Default: false
args-on-sep-lines: true

spancheck:
# Checks to enable.
# Options include:
# - end: check that span.End() is called
# - record-error: check that span.RecordError(err) is called when an error is returned
# - set-status: check that span.SetStatus(codes.Error, msg) is called when an error is returned
# Default: ["end"]
checks:
- end
- record-error
- set-status
# A list of regexes for function signatures that silence record-error and set-status reports
# if found in the call path to a returned error.
# https://github.com/jjti/go-spancheck#ignore-check-signatures
# Default: []
ignore-check-signatures:
- "telemetry.RecordError"
# A list of regexes for additional function signatures that create spans.
# This is useful if you have a utility method to create spans.
# Each entry should be of the form <regex>:<telemetry-type>, where telemetry-type can be opentelemetry or opencensus.
# https://github.com/jjti/go-spancheck#extra-start-span-signatures
# Default: []
extra-start-span-signatures:
- "github.com/user/repo/telemetry/trace.Start:opentelemetry"

staticcheck:
# https://staticcheck.dev/docs/configuration/options/#dot_import_whitelist
# Default: ["github.com/mmcloughlin/avo/build", "github.com/mmcloughlin/avo/operand", "github.com/mmcloughlin/avo/reg"]
dot-import-whitelist:
- fmt
# https://staticcheck.dev/docs/configuration/options/#initialisms
# Default: ["ACL", "API", "ASCII", "CPU", "CSS", "DNS", "EOF", "GUID", "HTML", "HTTP", "HTTPS", "ID", "IP", "JSON", "QPS", "RAM", "RPC", "SLA", "SMTP", "SQL", "SSH", "TCP", "TLS", "TTL", "UDP", "UI", "GID", "UID", "UUID", "URI", "URL", "UTF8", "VM", "XML", "XMPP", "XSRF", "XSS", "SIP", "RTP", "AMQP", "DB", "TS"]
initialisms: [ "ACL", "API", "ASCII", "CPU", "CSS", "DNS", "EOF", "GUID", "HTML", "HTTP", "HTTPS", "ID", "IP", "JSON", "QPS", "RAM", "RPC", "SLA", "SMTP", "SQL", "SSH", "TCP", "TLS", "TTL", "UDP", "UI", "GID", "UID", "UUID", "URI", "URL", "UTF8", "VM", "XML", "XMPP", "XSRF", "XSS", "SIP", "RTP", "AMQP", "DB", "TS" ]
# https://staticcheck.dev/docs/configuration/options/#http_status_code_whitelist
# Default: ["200", "400", "404", "500"]
http-status-code-whitelist: [ "200", "400", "404", "500" ]
# SAxxxx checks in https://staticcheck.dev/docs/configuration/options/#checks
# Example (to disable some checks): [ "all", "-SA1000", "-SA1001"]
# Default: ["all", "-ST1000", "-ST1003", "-ST1016", "-ST1020", "-ST1021", "-ST1022"]
checks:
# Invalid regular expression.
# https://staticcheck.dev/docs/checks/#SA1000
- SA1000
# Invalid template.
# https://staticcheck.dev/docs/checks/#SA1001
- SA1001
# Invalid format in 'time.Parse'.
# https://staticcheck.dev/docs/checks/#SA1002
- SA1002
# Unsupported argument to functions in 'encoding/binary'.
# https://staticcheck.dev/docs/checks/#SA1003
- SA1003
# Suspiciously small untyped constant in 'time.Sleep'.
# https://staticcheck.dev/docs/checks/#SA1004
- SA1004
# Invalid first argument to 'exec.Command'.
# https://staticcheck.dev/docs/checks/#SA1005
- SA1005
# 'Printf' with dynamic first argument and no further arguments.
# https://staticcheck.dev/docs/checks/#SA1006
- SA1006
# Invalid URL in 'net/url.Parse'.
# https://staticcheck.dev/docs/checks/#SA1007
- SA1007
# Non-canonical key in 'http.Header' map.
# https://staticcheck.dev/docs/checks/#SA1008
- SA1008
# '(*regexp.Regexp).FindAll' called with 'n == 0', which will always return zero results.
# https://staticcheck.dev/docs/checks/#SA1010
- SA1010
# Various methods in the "strings" package expect valid UTF-8, but invalid input is provided.
# https://staticcheck.dev/docs/checks/#SA1011
- SA1011
# A nil 'context.Context' is being passed to a function, consider using 'context.TODO' instead.
# https://staticcheck.dev/docs/checks/#SA1012
- SA1012
# 'io.Seeker.Seek' is being called with the whence constant as the first argument, but it should be the second.
# https://staticcheck.dev/docs/checks/#SA1013
- SA1013
# Non-pointer value passed to 'Unmarshal' or 'Decode'.
# https://staticcheck.dev/docs/checks/#SA1014
- SA1014
# Using 'time.Tick' in a way that will leak. Consider using 'time.NewTicker', and only use 'time.Tick' in tests, commands and endless functions.
# https://staticcheck.dev/docs/checks/#SA1015
- SA1015
# Trapping a signal that cannot be trapped.
# https://staticcheck.dev/docs/checks/#SA1016
- SA1016
# Channels used with 'os/signal.Notify' should be buffered.
# https://staticcheck.dev/docs/checks/#SA1017
- SA1017
# 'strings.Replace' called with 'n == 0', which does nothing.
# https://staticcheck.dev/docs/checks/#SA1018
- SA1018
# Using a deprecated function, variable, constant or field.
# https://staticcheck.dev/docs/checks/#SA1019
- SA1019
# Using an invalid host:port pair with a 'net.Listen'-related function.
# https://staticcheck.dev/docs/checks/#SA1020
- SA1020
# Using 'bytes.Equal' to compare two 'net.IP'.
# https://staticcheck.dev/docs/checks/#SA1021
- SA1021
# Modifying the buffer in an 'io.Writer' implementation.
# https://staticcheck.dev/docs/checks/#SA1023
- SA1023
# A string cutset contains duplicate characters.
# https://staticcheck.dev/docs/checks/#SA1024
- SA1024
# It is not possible to use '(*time.Timer).Reset''s return value correctly.
# https://staticcheck.dev/docs/checks/#SA1025
- SA1025
# Cannot marshal channels or functions.
# https://staticcheck.dev/docs/checks/#SA1026
- SA1026
# Atomic access to 64-bit variable must be 64-bit aligned.
# https://staticcheck.dev/docs/checks/#SA1027
- SA1027
# 'sort.Slice' can only be used on slices.
# https://staticcheck.dev/docs/checks/#SA1028
- SA1028
# Inappropriate key in call to 'context.WithValue'.
# https://staticcheck.dev/docs/checks/#SA1029
- SA1029
# Invalid argument in call to a 'strconv' function.
# https://staticcheck.dev/docs/checks/#SA1030
- SA1030
# Overlapping byte slices passed to an encoder.
# https://staticcheck.dev/docs/checks/#SA1031
- SA1031
# Wrong order of arguments to 'errors.Is'.
# https://staticcheck.dev/docs/checks/#SA1032
- SA1032
# 'sync.WaitGroup.Add' called inside the goroutine, leading to a race condition.
# https://staticcheck.dev/docs/checks/#SA2000
- SA2000
# Empty critical section, did you mean to defer the unlock?.
# https://staticcheck.dev/docs/checks/#SA2001
- SA2001
# Called 'testing.T.FailNow' or 'SkipNow' in a goroutine, which isn't allowed.
# https://staticcheck.dev/docs/checks/#SA2002
- SA2002
# Deferred 'Lock' right after locking, likely meant to defer 'Unlock' instead.
# https://staticcheck.dev/docs/checks/#SA2003
- SA2003
# 'TestMain' doesn't call 'os.Exit', hiding test failures.
# https://staticcheck.dev/docs/checks/#SA3000
- SA3000
# Assigning to 'b.N' in benchmarks distorts the results.
# https://staticcheck.dev/docs/checks/#SA3001
- SA3001
# Binary operator has identical expressions on both sides.
# https://staticcheck.dev/docs/checks/#SA4000
- SA4000
# '&*x' gets simplified to 'x', it does not copy 'x'.
# https://staticcheck.dev/docs/checks/#SA4001
- SA4001
# Comparing unsigned values against negative values is pointless.
# https://staticcheck.dev/docs/checks/#SA4003
- SA4003
# The loop exits unconditionally after one iteration.
# https://staticcheck.dev/docs/checks/#SA4004
- SA4004
# Field assignment that will never be observed. Did you mean to use a pointer receiver?.
# https://staticcheck.dev/docs/checks/#SA4005
- SA4005
# A value assigned to a variable is never read before being overwritten. Forgotten error check or dead code?.
# https://staticcheck.dev/docs/checks/#SA4006
- SA4006
# The variable in the loop condition never changes, are you incrementing the wrong variable?.
# https://staticcheck.dev/docs/checks/#SA4008
- SA4008
# A function argument is overwritten before its first use.
# https://staticcheck.dev/docs/checks/#SA4009
- SA4009
# The result of 'append' will never be observed anywhere.
# https://staticcheck.dev/docs/checks/#SA4010
- SA4010
# Break statement with no effect. Did you mean to break out of an outer loop?.
# https://staticcheck.dev/docs/checks/#SA4011
- SA4011
# Comparing a value against NaN even though no value is equal to NaN.
# https://staticcheck.dev/docs/checks/#SA4012
- SA4012
# Negating a boolean twice ('!!b') is the same as writing 'b'. This is either redundant, or a typo.
# https://staticcheck.dev/docs/checks/#SA4013
- SA4013
# An if/else if chain has repeated conditions and no side-effects; if the condition didn't match the first time, it won't match the second time, either.
# https://staticcheck.dev/docs/checks/#SA4014
- SA4014
# Calling functions like 'math.Ceil' on floats converted from integers doesn't do anything useful.
# https://staticcheck.dev/docs/checks/#SA4015
- SA4015
# Certain bitwise operations, such as 'x ^ 0', do not do anything useful.
# https://staticcheck.dev/docs/checks/#SA4016
- SA4016
# Discarding the return values of a function without side effects, making the call pointless.
# https://staticcheck.dev/docs/checks/#SA4017
- SA4017
# Self-assignment of variables.
# https://staticcheck.dev/docs/checks/#SA4018
- SA4018
# Multiple, identical build constraints in the same file.
# https://staticcheck.dev/docs/checks/#SA4019
- SA4019
# Unreachable case clause in a type switch.
# https://staticcheck.dev/docs/checks/#SA4020
- SA4020
# "x = append(y)" is equivalent to "x = y".
# https://staticcheck.dev/docs/checks/#SA4021
- SA4021
# Comparing the address of a variable against nil.
# https://staticcheck.dev/docs/checks/#SA4022
- SA4022
# Impossible comparison of interface value with untyped nil.
# https://staticcheck.dev/docs/checks/#SA4023
- SA4023
# Checking for impossible return value from a builtin function.
# https://staticcheck.dev/docs/checks/#SA4024
- SA4024
# Integer division of literals that results in zero.
# https://staticcheck.dev/docs/checks/#SA4025
- SA4025
# Go constants cannot express negative zero.
# https://staticcheck.dev/docs/checks/#SA4026
- SA4026
# '(*net/url.URL).Query' returns a copy, modifying it doesn't change the URL.
# https://staticcheck.dev/docs/checks/#SA4027
- SA4027
# 'x % 1' is always zero.
# https://staticcheck.dev/docs/checks/#SA4028
- SA4028
# Ineffective attempt at sorting slice.
# https://staticcheck.dev/docs/checks/#SA4029
- SA4029
# Ineffective attempt at generating random number.
# https://staticcheck.dev/docs/checks/#SA4030
- SA4030
# Checking never-nil value against nil.
# https://staticcheck.dev/docs/checks/#SA4031
- SA4031
# Comparing 'runtime.GOOS' or 'runtime.GOARCH' against impossible value.
# https://staticcheck.dev/docs/checks/#SA4032
- SA4032
# Assignment to nil map.
# https://staticcheck.dev/docs/checks/#SA5000
- SA5000
# Deferring 'Close' before checking for a possible error.
# https://staticcheck.dev/docs/checks/#SA5001
- SA5001
# The empty for loop ("for {}") spins and can block the scheduler.
# https://staticcheck.dev/docs/checks/#SA5002
- SA5002
# Defers in infinite loops will never execute.
# https://staticcheck.dev/docs/checks/#SA5003
- SA5003
# "for { select { ..." with an empty default branch spins.
# https://staticcheck.dev/docs/checks/#SA5004
- SA5004
# The finalizer references the finalized object, preventing garbage collection.
# https://staticcheck.dev/docs/checks/#SA5005
- SA5005
# Infinite recursive call.
# https://staticcheck.dev/docs/checks/#SA5007
- SA5007
# Invalid struct tag.
# https://staticcheck.dev/docs/checks/#SA5008
- SA5008
# Invalid Printf call.
# https://staticcheck.dev/docs/checks/#SA5009
- SA5009
# Impossible type assertion.
# https://staticcheck.dev/docs/checks/#SA5010
- SA5010
# Possible nil pointer dereference.
# https://staticcheck.dev/docs/checks/#SA5011
- SA5011
# Passing odd-sized slice to function expecting even size.
# https://staticcheck.dev/docs/checks/#SA5012
- SA5012
# Using 'regexp.Match' or related in a loop, should use 'regexp.Compile'.
# https://staticcheck.dev/docs/checks/#SA6000
- SA6000
# Missing an optimization opportunity when indexing maps by byte slices.
# https://staticcheck.dev/docs/checks/#SA6001
- SA6001
# Storing non-pointer values in 'sync.Pool' allocates memory.
# https://staticcheck.dev/docs/checks/#SA6002
- SA6002
# Converting a string to a slice of runes before ranging over it.
# https://staticcheck.dev/docs/checks/#SA6003
- SA6003
# Inefficient string comparison with 'strings.ToLower' or 'strings.ToUpper'.
# https://staticcheck.dev/docs/checks/#SA6005
- SA6005
# Using io.WriteString to write '[]byte'.
# https://staticcheck.dev/docs/checks/#SA6006
- SA6006
# Defers in range loops may not run when you expect them to.
# https://staticcheck.dev/docs/checks/#SA9001
- SA9001
# Using a non-octal 'os.FileMode' that looks like it was meant to be in octal.
# https://staticcheck.dev/docs/checks/#SA9002
- SA9002
# Empty body in an if or else branch.
# https://staticcheck.dev/docs/checks/#SA9003
- SA9003
# Only the first constant has an explicit type.
# https://staticcheck.dev/docs/checks/#SA9004
- SA9004
# Trying to marshal a struct with no public fields nor custom marshaling.
# https://staticcheck.dev/docs/checks/#SA9005
- SA9005
# Dubious bit shifting of a fixed size integer value.
# https://staticcheck.dev/docs/checks/#SA9006
- SA9006
# Deleting a directory that shouldn't be deleted.
# https://staticcheck.dev/docs/checks/#SA9007
- SA9007
# 'else' branch of a type assertion is probably not reading the right value.
# https://staticcheck.dev/docs/checks/#SA9008
- SA9008
# Ineffectual Go compiler directive.
# https://staticcheck.dev/docs/checks/#SA9009
- SA9009
# Incorrect or missing package comment.
# https://staticcheck.dev/docs/checks/#ST1000
- ST1000
# Dot imports are discouraged.
# https://staticcheck.dev/docs/checks/#ST1001
- ST1001
# Poorly chosen identifier.
# https://staticcheck.dev/docs/checks/#ST1003
- ST1003
# Incorrectly formatted error string.
# https://staticcheck.dev/docs/checks/#ST1005
- ST1005
# Poorly chosen receiver name.
# https://staticcheck.dev/docs/checks/#ST1006
- ST1006
# A function's error value should be its last return value.
# https://staticcheck.dev/docs/checks/#ST1008
- ST1008
# Poorly chosen name for variable of type 'time.Duration'.
# https://staticcheck.dev/docs/checks/#ST1011
- ST1011
# Poorly chosen name for error variable.
# https://staticcheck.dev/docs/checks/#ST1012
- ST1012
# Should use constants for HTTP error codes, not magic numbers.
# https://staticcheck.dev/docs/checks/#ST1013
- ST1013
# A switch's default case should be the first or last case.
# https://staticcheck.dev/docs/checks/#ST1015
- ST1015
# Use consistent method receiver names.
# https://staticcheck.dev/docs/checks/#ST1016
- ST1016
# Don't use Yoda conditions.
# https://staticcheck.dev/docs/checks/#ST1017
- ST1017
# Avoid zero-width and control characters in string literals.
# https://staticcheck.dev/docs/checks/#ST1018
- ST1018
# Importing the same package multiple times.
# https://staticcheck.dev/docs/checks/#ST1019
- ST1019
# The documentation of an exported function should start with the function's name.
# https://staticcheck.dev/docs/checks/#ST1020
- ST1020
# The documentation of an exported type should start with type's name.
# https://staticcheck.dev/docs/checks/#ST1021
- ST1021
# The documentation of an exported variable or constant should start with variable's name.
# https://staticcheck.dev/docs/checks/#ST1022
- ST1022
# Redundant type in variable declaration.
# https://staticcheck.dev/docs/checks/#ST1023
- ST1023
# Use plain channel send or receive instead of single-case select.
# https://staticcheck.dev/docs/checks/#S1000
- S1000
# Replace for loop with call to copy.
# https://staticcheck.dev/docs/checks/#S1001
- S1001
# Omit comparison with boolean constant.
# https://staticcheck.dev/docs/checks/#S1002
- S1002
# Replace call to 'strings.Index' with 'strings.Contains'.
# https://staticcheck.dev/docs/checks/#S1003
- S1003
# Replace call to 'bytes.Compare' with 'bytes.Equal'.
# https://staticcheck.dev/docs/checks/#S1004
- S1004
# Drop unnecessary use of the blank identifier.
# https://staticcheck.dev/docs/checks/#S1005
- S1005
# Use "for { ... }" for infinite loops.
# https://staticcheck.dev/docs/checks/#S1006
- S1006
# Simplify regular expression by using raw string literal.
# https://staticcheck.dev/docs/checks/#S1007
- S1007
# Simplify returning boolean expression.
# https://staticcheck.dev/docs/checks/#S1008
- S1008
# Omit redundant nil check on slices, maps, and channels.
# https://staticcheck.dev/docs/checks/#S1009
- S1009
# Omit default slice index.
# https://staticcheck.dev/docs/checks/#S1010
- S1010
# Use a single 'append' to concatenate two slices.
# https://staticcheck.dev/docs/checks/#S1011
- S1011
# Replace 'time.Now().Sub(x)' with 'time.Since(x)'.
# https://staticcheck.dev/docs/checks/#S1012
- S1012
# Use a type conversion instead of manually copying struct fields.
# https://staticcheck.dev/docs/checks/#S1016
- S1016
# Replace manual trimming with 'strings.TrimPrefix'.
# https://staticcheck.dev/docs/checks/#S1017
- S1017
# Use "copy" for sliding elements.
# https://staticcheck.dev/docs/checks/#S1018
- S1018
# Simplify "make" call by omitting redundant arguments.
# https://staticcheck.dev/docs/checks/#S1019
- S1019
# Omit redundant nil check in type assertion.
# https://staticcheck.dev/docs/checks/#S1020
- S1020
# Merge variable declaration and assignment.
# https://staticcheck.dev/docs/checks/#S1021
- S1021
# Omit redundant control flow.
# https://staticcheck.dev/docs/checks/#S1023
- S1023
# Replace 'x.Sub(time.Now())' with 'time.Until(x)'.
# https://staticcheck.dev/docs/checks/#S1024
- S1024
# Don't use 'fmt.Sprintf("%s", x)' unnecessarily.
# https://staticcheck.dev/docs/checks/#S1025
- S1025
# Simplify error construction with 'fmt.Errorf'.
# https://staticcheck.dev/docs/checks/#S1028
- S1028
# Range over the string directly.
# https://staticcheck.dev/docs/checks/#S1029
- S1029
# Use 'bytes.Buffer.String' or 'bytes.Buffer.Bytes'.
# https://staticcheck.dev/docs/checks/#S1030
- S1030
# Omit redundant nil check around loop.
# https://staticcheck.dev/docs/checks/#S1031
- S1031
# Use 'sort.Ints(x)', 'sort.Float64s(x)', and 'sort.Strings(x)'.
# https://staticcheck.dev/docs/checks/#S1032
- S1032
# Unnecessary guard around call to "delete".
# https://staticcheck.dev/docs/checks/#S1033
- S1033
# Use result of type assertion to simplify cases.
# https://staticcheck.dev/docs/checks/#S1034
- S1034
# Redundant call to 'net/http.CanonicalHeaderKey' in method call on 'net/http.Header'.
# https://staticcheck.dev/docs/checks/#S1035
- S1035
# Unnecessary guard around map access.
# https://staticcheck.dev/docs/checks/#S1036
- S1036
# Elaborate way of sleeping.
# https://staticcheck.dev/docs/checks/#S1037
- S1037
# Unnecessarily complex way of printing formatted string.
# https://staticcheck.dev/docs/checks/#S1038
- S1038
# Unnecessary use of 'fmt.Sprint'.
# https://staticcheck.dev/docs/checks/#S1039
- S1039
# Type assertion to current type.
# https://staticcheck.dev/docs/checks/#S1040
- S1040
# Apply De Morgan's law.
# https://staticcheck.dev/docs/checks/#QF1001
- QF1001
# Convert untagged switch to tagged switch.
# https://staticcheck.dev/docs/checks/#QF1002
- QF1002
# Convert if/else-if chain to tagged switch.
# https://staticcheck.dev/docs/checks/#QF1003
- QF1003
# Use 'strings.ReplaceAll' instead of 'strings.Replace' with 'n == -1'.
# https://staticcheck.dev/docs/checks/#QF1004
- QF1004
# Expand call to 'math.Pow'.
# https://staticcheck.dev/docs/checks/#QF1005
- QF1005
# Lift 'if'+'break' into loop condition.
# https://staticcheck.dev/docs/checks/#QF1006
- QF1006
# Merge conditional assignment into variable declaration.
# https://staticcheck.dev/docs/checks/#QF1007
- QF1007
# Omit embedded fields from selector expression.
# https://staticcheck.dev/docs/checks/#QF1008
- QF1008
# Use 'time.Time.Equal' instead of '==' operator.
# https://staticcheck.dev/docs/checks/#QF1009
- QF1009
# Convert slice of bytes to string when printing it.
# https://staticcheck.dev/docs/checks/#QF1010
- QF1010
# Omit redundant type from variable declaration.
# https://staticcheck.dev/docs/checks/#QF1011
- QF1011
# Use 'fmt.Fprintf(x, ...)' instead of 'x.Write(fmt.Sprintf(...))'.
# https://staticcheck.dev/docs/checks/#QF1012
- QF1012

tagalign:
# Align and sort can be used together or separately.
#
# Whether enable align. If true, the struct tags will be aligned.
# e.g.:
# type FooBar struct {
	#     Bar    string json:"bar" validate:"required"
	#     FooFoo int8   json:"foo_foo" validate:"required"
	# }
# will be formatted to:
# type FooBar struct {
	#     Bar    string json:"bar"     validate:"required"
	#     FooFoo int8   json:"foo_foo" validate:"required"
	# }
# Default: true.
align: false
# Whether enable tags sort.
# If true, the tags will be sorted by name in ascending order.
# e.g.: xml:"bar" json:"bar" validate:"required" -> json:"bar" validate:"required" xml:"bar"
# Default: true
sort: false
# Specify the order of tags, the other tags will be sorted by name.
# This option will be ignored if sort is false.
# Default: []
order:
- json
- yaml
- yml
- toml
- mapstructure
- binding
- validate
# Whether enable strict style.
# In this style, the tags will be sorted and aligned in the dictionary order,
# and the tags with the same name will be aligned together.
# Note: This option will be ignored if 'align' or 'sort' is false.
# Default: false
strict: true

tagliatelle:
# Checks the struct tag name case.
case:
# Defines the association between tag name and case.
# Any struct tag name can be used.
# Supported string cases:
# - camel
# - pascal
# - kebab
# - snake
# - upperSnake
# - goCamel
# - goPascal
# - goKebab
# - goSnake
# - upper
# - lower
# - header
rules:
json: camel
yaml: camel
xml: camel
toml: camel
bson: camel
avro: snake
mapstructure: kebab
env: upperSnake
envconfig: upperSnake
whatever: snake
# Defines the association between tag name and case.
# Important: the extended-rules overrides rules.
# Default: empty
extended-rules:
json:
# Supported string cases:
# - camel
# - pascal
# - kebab
# - snake
# - upperSnake
# - goCamel
# - goPascal
# - goKebab
# - goSnake
# - header
# - lower
# - header
#
# Required
case: camel
# Adds 'AMQP', 'DB', 'GID', 'RTP', 'SIP', 'TS' to initialisms,
# and removes 'LHS', 'RHS' from initialisms.
# Default: false
extra-initialisms: true
# Defines initialism additions and overrides.
# Default: empty
initialism-overrides:
DB: true # add a new initialism
LHS: false # disable a default initialism.
# ...
# Uses the struct field name to check the name of the struct tag.
# Default: false
use-field-name: true
# The field names to ignore.
# Default: []
ignored-fields:
- Bar
- Foo
# Overrides the default/root configuration.
# Default: []
overrides:
-
# The package path (uses / only as a separator).
# Required
pkg: foo/bar
# Default: empty or the same as the default/root configuration.
rules:
json: snake
xml: pascal
# Default: empty or the same as the default/root configuration.
extended-rules:
# Same options as the base extended-rules.
# Default: false (WARNING: it doesn't follow the default/root configuration)
use-field-name: true
# The field names to ignore.
# Default: [] or the same as the default/root configuration.
ignored-fields:
- Bar
- Foo
# Ignore the package (takes precedence over all other configurations).
# Default: false
ignore: true

tenv:
# The option all will run against whole test files (_test.go) regardless of method/function signatures.
# Otherwise, only methods that take *testing.T, *testing.B, and testing.TB as arguments are checked.
# Default: false
all: false

testifylint:
# Enable all checkers (https://github.com/Antonboom/testifylint#checkers).
# Default: false
enable-all: true
# Disable checkers by name
# (in addition to default
#   suite-thelper
# ).
disable:
- blank-import
- bool-compare
- compares
- contains
- empty
- encoded-compare
- equal-values
- error-is-as
- error-nil
- expected-actual
- float-compare
- formatter
- go-require
- len
- negative-positive
- nil-compare
- regexp
- require-error
- suite-broken-parallel
- suite-dont-use-pkg
- suite-extra-assert-call
- suite-method-signature
- suite-subtest-run
- suite-thelper
- useless-assert

# Disable all checkers (https://github.com/Antonboom/testifylint#checkers).
# Default: false
disable-all: true
# Enable checkers by name
# (in addition to default
#   blank-import, bool-compare, compares, contains, empty, encoded-compare, equal-values, error-is-as, error-nil,
#   expected-actual, go-require, float-compare, formatter, len, negative-positive, nil-compare, regexp, require-error,
#   suite-broken-parallel, suite-dont-use-pkg, suite-extra-assert-call, suite-subtest-run, suite-method-signature,
#   useless-assert
# ).
enable:
- blank-import
- bool-compare
- compares
- contains
- empty
- encoded-compare
- equal-values
- error-is-as
- error-nil
- expected-actual
- float-compare
- formatter
- go-require
- len
- negative-positive
- nil-compare
- regexp
- require-error
- suite-broken-parallel
- suite-dont-use-pkg
- suite-extra-assert-call
- suite-method-signature
- suite-subtest-run
- suite-thelper
- useless-assert

bool-compare:
# To ignore user defined types (over builtin bool).
# Default: false
ignore-custom-types: true
expected-actual:
# Regexp for expected variable name.
# Default: (^(exp(ected)?|want(ed)?)([A-Z]\w*)?$)|(^(\w*[a-z])?(Exp(ected)?|Want(ed)?)$)
pattern: ^expected
formatter:
# To enable go vet's printf checks.
# Default: true
check-format-string: false
# To require f-assertions (e.g. assert.Equalf) if format string is used, even if there are no variable-length
# variables, i.e. it requires require.NoErrorf for both these cases:
# - require.NoErrorf(t, err, "unexpected error")
# - require.NoErrorf(t, err, "unexpected error for sid: %v", sid)
# To understand this behavior, please read the
# https://github.com/Antonboom/testifylint?tab=readme-ov-file#historical-reference-of-formatter.
# Default: false
require-f-funcs: true
# To require that the first element of msgAndArgs (msg) has a string type.
# For example, in such case assertion like assert.True(t, b, tt.case) will be considered as invalid.
# Default: true
require-string-msg: false
go-require:
# To ignore HTTP handlers (like http.HandlerFunc).
# Default: false
ignore-http-handlers: true
require-error:
# Regexp for assertions to analyze. If defined, then only matched error assertions will be reported.
# Default: ""
fn-pattern: ^(Errorf?|NoErrorf?)$
suite-extra-assert-call:
# To require or remove extra Assert() call?
# Default: remove
mode: require

testpackage:
# Regexp pattern to skip files.
# Default: "(export|internal)_test\\.go"
skip-regexp: (export|internal)_test\.go
# List of packages that don't end with _test that tests are allowed to be in.
# Default: "main"
allow-packages:
- example
- main

usestdlibvars:
# Suggest the use of http.MethodXX.
# Default: true
http-method: false
# Suggest the use of http.StatusXX.
# Default: true
http-status-code: false
# Suggest the use of time.Weekday.String().
# Default: true
time-weekday: true
# Suggest the use of time.Month.String().
# Default: false
time-month: true
# Suggest the use of time.Layout.
# Default: false
time-layout: true
# Suggest the use of crypto.Hash.String().
# Default: false
crypto-hash: true
# Suggest the use of rpc.DefaultXXPath.
# Default: false
default-rpc-path: true
# Suggest the use of sql.LevelXX.String().
# Default: false
sql-isolation-level: true
# Suggest the use of tls.SignatureScheme.String().
# Default: false
tls-signature-scheme: true
# Suggest the use of constant.Kind.String().
# Default: false
constant-kind: true

usetesting:
# Enable/disable os.CreateTemp("", ...) detections.
# Default: true
os-create-temp: false

# Enable/disable os.MkdirTemp() detections.
# Default: true
os-mkdir-temp: false

# Enable/disable os.Setenv() detections.
# Default: true
os-setenv: false

# Enable/disable os.TempDir() detections.
# Default: false
os-temp-dir: true

# Enable/disable os.Chdir() detections.
# Disabled if Go < 1.24.
# Default: true
os-chdir: false

# Enable/disable context.Background() detections.
# Disabled if Go < 1.24.
# Default: true
context-background: false

# Enable/disable context.TODO() detections.
# Disabled if Go < 1.24.
# Default: true
context-todo: false

unconvert:
# Remove conversions that force intermediate rounding.
# Default: false
fast-math: true
# Be more conservative (experimental).
# Default: false
safe: true

unparam:
# Inspect exported functions.
#
# Set to true if no external program/library imports your code.
# XXX: if you enable this setting, unparam will report a lot of false-positives in text editors:
# if it's called for subdir of a project it can't find external interfaces. All text editor integrations
# with golangci-lint call it on a directory with the changed file.
#
# Default: false
check-exported: true

unused:
# Mark all struct fields that have been written to as used.
# Default: true
field-writes-are-uses: false
# Treat IncDec statement (e.g. i++ or i--) as both read and write operation instead of just write.
# Default: false
post-statements-are-reads: true
# Mark all exported fields as used.
# default: true
exported-fields-are-used: false
# Mark all function parameters as used.
# default: true
parameters-are-used: false
# Mark all local variables as used.
# default: true
local-variables-are-used: false
# Mark all identifiers inside generated files as used.
# Default: true
generated-is-used: false

varnamelen:
# The longest distance, in source lines, that is being considered a "small scope".
# Variables used in at most this many lines will be ignored.
# Default: 5
max-distance: 6
# The minimum length of a variable's name that is considered "long".
# Variable names that are at least this long will be ignored.
# Default: 3
min-name-length: 2
# Check method receivers.
# Default: false
check-receiver: true
# Check named return values.
# Default: false
check-return: true
# Check type parameters.
# Default: false
check-type-param: true
# Ignore "ok" variables that hold the bool return value of a type assertion.
# Default: false
ignore-type-assert-ok: true
# Ignore "ok" variables that hold the bool return value of a map index.
# Default: false
ignore-map-index-ok: true
# Ignore "ok" variables that hold the bool return value of a channel receive.
# Default: false
ignore-chan-recv-ok: true
# Optional list of variable names that should be ignored completely.
# Default: []
ignore-names:
- err
# Optional list of variable declarations that should be ignored completely.
# Entries must be in one of the following forms (see below for examples):
# - for variables, parameters, named return values, method receivers, or type parameters:
#   <name> <type>  (<type> can also be a pointer/slice/map/chan/...)
# - for constants: const <name>
#
# Default: []
ignore-decls:
- c echo.Context
- t testing.T
- f *foo.Bar
- e error
- i int
- const C
- T any
- m map[string]int

whitespace:
# Enforces newlines (or comments) after every multi-line if statement.
# Default: false
multi-if: true
# Enforces newlines (or comments) after every multi-line function signature.
# Default: false
multi-func: true

wrapcheck:
# An array of strings specifying additional substrings of signatures to ignore.
# Unlike 'ignore-sigs', this option extends the default set (or the set specified in 'ignore-sigs') without replacing it entirely.
# This allows you to add specific signatures to the ignore list
# while retaining the defaults or any items in 'ignore-sigs'.
# Default: []
extra-ignore-sigs:
- .CustomError(
- .SpecificWrap(

# An array of strings that specify substrings of signatures to ignore.
# If this set, it will override the default set of ignored signatures.
# See https://github.com/tomarrell/wrapcheck#configuration for more information.
# Default: [".Errorf(", "errors.New(", "errors.Unwrap(", "errors.Join(", ".Wrap(", ".Wrapf(", ".WithMessage(", ".WithMessagef(", ".WithStack("]
ignore-sigs:
- .Errorf(
- errors.New(
- errors.Unwrap(
- errors.Join(
- .Wrap(
- .Wrapf(
- .WithMessage(
- .WithMessagef(
- .WithStack(
# An array of strings that specify regular expressions of signatures to ignore.
# Default: []
ignore-sig-regexps:
- \.New.*Error\(
# An array of strings that specify globs of packages to ignore.
# Default: []
ignore-package-globs:
- encoding/*
        - github.com/pkg/*
      # An array of strings that specify regular expressions of interfaces to ignore.
      # Default: []
      ignore-interface-regexps:
        - ^(?i)c(?-i)ach(ing|e)

    wsl:
      # Do strict checking when assigning from append (x = append(x, y)).
      # If this is set to true - the append call must append either a variable
      # assigned, called or used on the line above.
      # https://github.com/bombsimon/wsl/blob/HEAD/doc/configuration.md#strict-append
      # Default: true
      strict-append: false

      # Allows assignments to be cuddled with variables used in calls on
      # line above and calls to be cuddled with assignments of variables
      # used in call on line above.
      # https://github.com/bombsimon/wsl/blob/HEAD/doc/configuration.md#allow-assign-and-call
      # Default: true
      allow-assign-and-call: false

      # Allows assignments to be cuddled with anything.
      # https://github.com/bombsimon/wsl/blob/HEAD/doc/configuration.md#allow-assign-and-anything
      # Default: false
      allow-assign-and-anything: true

      # Allows cuddling to assignments even if they span over multiple lines.
      # https://github.com/bombsimon/wsl/blob/HEAD/doc/configuration.md#allow-multiline-assign
      # Default: true
      allow-multiline-assign: false

      # If the number of lines in a case block is equal to or lager than this number,
      # the case *must* end white a newline.
      # https://github.com/bombsimon/wsl/blob/HEAD/doc/configuration.md#force-case-trailing-whitespace
      # Default: 0
      force-case-trailing-whitespace: 1

      # Allow blocks to end with comments.
      # https://github.com/bombsimon/wsl/blob/HEAD/doc/configuration.md#allow-trailing-comment
      # Default: false
      allow-trailing-comment: true

      # Allow multiple comments in the beginning of a block separated with newline.
      # https://github.com/bombsimon/wsl/blob/HEAD/doc/configuration.md#allow-separated-leading-comment
      # Default: false
      allow-separated-leading-comment: true

      # Allow multiple var/declaration statements to be cuddled.
      # https://github.com/bombsimon/wsl/blob/HEAD/doc/configuration.md#allow-cuddle-declarations
      # Default: false
      allow-cuddle-declarations: true

      # A list of call idents that everything can be cuddled with.
      # Defaults: [ "Lock", "RLock" ]
      allow-cuddle-with-calls: [ "Foo", "Bar" ]

      # AllowCuddleWithRHS is a list of right hand side variables that is allowed
      # to be cuddled with anything.
      # Defaults: [ "Unlock", "RUnlock" ]
      allow-cuddle-with-rhs: [ "Foo", "Bar" ]

      # Allow cuddling with any block as long as the variable is used somewhere in
      # the block.
      # https://github.com/bombsimon/wsl/blob/HEAD/doc/configuration.md#allow-cuddle-used-in-block
      # Default: false
      allow-cuddle-used-in-block: true

      # Causes an error when an If statement that checks an error variable doesn't
      # cuddle with the assignment of that variable.
      # https://github.com/bombsimon/wsl/blob/HEAD/doc/configuration.md#force-err-cuddling
      # Default: false
      force-err-cuddling: true

      # When force-err-cuddling is enabled this is a list of names
      # used for error variables to check for in the conditional.
      # Default: [ "err" ]
      error-variable-names: [ "foo" ]

      # Causes an error if a short declaration (:=) cuddles with anything other than
      # another short declaration.
      # This logic overrides force-err-cuddling among others.
      # https://github.com/bombsimon/wsl/blob/HEAD/doc/configuration.md#force-short-decl-cuddling
      # Default: false
      force-short-decl-cuddling: true

  # Defines a set of rules to ignore issues.
  # It does not skip the analysis, and so does not ignore "typecheck" errors.
  exclusions:
    # Mode of the generated files analysis.
    #
    # - strict: sources are excluded by strictly following the Go generated file convention.
    #    Source files that have lines matching only the following regular expression will be excluded: ^// Code generated .* DO NOT EDIT\.$
    #    This line must appear before the first non-comment, non-blank text in the file.
    #    https://go.dev/s/generatedcode
    # - lax: sources are excluded if they contain lines like autogenerated file, code generated, do not edit, etc.
    # - disable: disable the generated files exclusion.
    #
    # Default: lax
    generated: strict
    # Log a warning if an exclusion rule is unused.
    # Default: false
    warn-unused: true
    # Predefined exclusion rules.
    # Default: []
    presets:
      - comments
      - std-error-handling
      - common-false-positives
      - legacy

    # Excluding configuration per-path, per-linter, per-text and per-source.
    rules:
      # Exclude some linters from running on tests files.
      - path: "internal/controllers/grpc*"
        linters:
          - revive
      - path: "_test\\.go"
        linters:
          - bodyclose
          - dupl
          - funlen
          - goconst
          - gosec
          - noctx
          - wrapcheck
      - path: "cmd/client/client.go"
        linters:
          - mnd
          - govet
      - path: "internal/config/config.go"
        linters:
          - nosprintfhostport
          - funlen

      # Run some linter only for test files by excluding its issues for everything else.
      - path-except: _test\.go
        linters:
          - forbidigo

      # Exclude known linters from partially hard-vendored code,
      # which is impossible to exclude via nolint comments.
      # / will be replaced by the current OS file path separator to properly work on Windows.
      - path: internal/hmac/
        text: "weak cryptographic primitive"
        linters:
          - gosec

      # Exclude some staticcheck messages.
      - linters:
          - staticcheck
        text: "SA9003:"

      # Exclude lll issues for long lines with go:generate.
      - linters:
          - lll
        source: "^//go:generate "

    # Which file paths to exclude: they will be analyzed, but issues from them won't be reported.
    # "/" will be replaced by the current OS file path separator to properly work on Windows.
    # Default: []
    paths:
      - ".*\\.my\\.go$"
      - lib/bad.go
    # Which file paths to not exclude.
    # Default: []
    paths-except:
      - ".*\\.my\\.go$"
      - lib/bad.go

formatters:
  # Enable specific formatter.
  # Default: [] (uses standard Go formatting)
  enable:
    - gci
    - gofmt
    - gofumpt
    - goimports
    - golines

  # Formatters settings.
  settings:
    gci:
      # Section configuration to compare against.
      # Section names are case-insensitive and may contain parameters in ().
      # The default order of sections is standard > default > custom > blank > dot > alias > localmodule,
      # If custom-order is true, it follows the order of sections option.
      # Default: ["standard", "default"]
      sections:
        - standard                       # Standard section: captures all standard packages.
        - default                        # Default section: contains all imports that could not be matched to another section type.
        - prefix(github.com/org/project) # Custom section: groups all imports with the specified Prefix.
        - blank                          # Blank section: contains all blank imports. This section is not present unless explicitly enabled.
        - dot                            # Dot section: contains all dot imports. This section is not present unless explicitly enabled.
        - alias                          # Alias section: contains all alias imports. This section is not present unless explicitly enabled.
        - localmodule                    # Local module section: contains all local packages. This section is not present unless explicitly enabled.

      # Checks that no inline comments are present.
      # Default: false
      no-inline-comments: true

      # Checks that no prefix comments (comment lines above an import) are present.
      # Default: false
      no-prefix-comments: true

      # Enable custom order of sections.
      # If true, make the section order the same as the order of sections.
      # Default: false
      custom-order: true

      # Drops lexical ordering for custom sections.
      # Default: false
      no-lex-order: true

    gofmt:
      # Simplify code: gofmt with -s option.
      # Default: true
      simplify: false
      # Apply the rewrite rules to the source before reformatting.
      # https://pkg.go.dev/cmd/gofmt
      # Default: []
      rewrite-rules:
        - pattern: 'interface{}'
          replacement: 'any'
        - pattern: 'a[b:len(a)]'
          replacement: 'a[b:]'

    gofumpt:
      # Module path which contains the source code being formatted.
      # Default: ""
      module-path: github.com/org/project

      # Choose whether to use the extra rules.
      # Default: false
      extra-rules: true

    goimports:
      # A list of prefixes, which, if set, checks import paths
      # with the given prefixes are grouped after 3rd-party packages.
      # Default: []
      local-prefixes:
        - github.com/org/project

    golines:
      # Target maximum line length.
      # Default: 100
      max-len: 200
      # Length of a tabulation.
      # Default: 4
      tab-len: 8
      # Shorten single-line comments.
      # Default: false
      shorten-comments: true
      # Default: true
      reformat-tags: false
      # Split chained methods on the dots as opposed to the arguments.
      # Default: true
      chain-split-dots: false

  exclusions:
    # Mode of the generated files analysis.
    #
    # - strict: sources are excluded by strictly following the Go generated file convention.
    #    Source files that have lines matching only the following regular expression will be excluded: ^// Code generated .* DO NOT EDIT\.$
    #    This line must appear before the first non-comment, non-blank text in the file.
    #    https://go.dev/s/generatedcode
    # - lax: sources are excluded if they contain lines like autogenerated file, code generated, do not edit, etc.
    # - disable: disable the generated files exclusion.
    #
    # Default: lax
    generated: strict
    # Which file paths to exclude.
    # Default: []
    paths:
      - ".*\\.my\\.go$"
      - lib/bad.go

#issues:
#  # Maximum issues count per one linter.
#  # Set to 0 to disable.
#  # Default: 50
#  max-issues-per-linter: 0
#
#  # Maximum count of issues with the same text.
#  # Set to 0 to disable.
#  # Default: 3
#  max-same-issues: 0
#
#  # Make issues output unique by line.
#  # Default: true
#  uniq-by-line: false
#
#  # Show only new issues: if there are unstaged changes or untracked files,
#  # only those changes are analyzed, else only changes in HEAD~ are analyzed.
#  # It's a super-useful option for integration of golangci-lint into existing large codebase.
#  # It's not practical to fix all existing issues at the moment of integration:
#  # much better don't allow issues in new code.
#  #
#  # Default: false
#  new: true
#
#  # Show only new issues created after the best common ancestor (merge-base against HEAD).
#  # Default: ""
#  new-from-merge-base: main
#
#  # Show only new issues created after git revision REV.
#  # Default: ""
#  new-from-rev: HEAD
#
#  # Show only new issues created in git patch with set file path.
#  # Default: ""
#  new-from-patch: path/to/patch/file
#
#  # Show issues in any part of update files (requires new-from-rev or new-from-patch).
#  # Default: false
#  whole-files: true
#
#  # Fix found issues (if it's supported by the linter).
#  # Default: false
#  fix: true


## Output configuration options.
#output:
#  # The formats used to render issues.
#  formats:
#    # Prints issues in a text format with colors, line number, and linter name.
#    # This format is the default format.
#    text:
#      # Output path can be either stdout, stderr or path to the file to write to.
#      # Default: stdout
#      path: ./path/to/output.txt
#      # Print linter name in the end of issue text.
#      # Default: true
#      print-linter-name: false
#      # Print lines of code with issue.
#      # Default: true
#      print-issued-lines: false
#      # Use colors.
#      # Default: true
#      colors: false
#    # Prints issues in a JSON representation.
#    json:
#      # Output path can be either stdout, stderr or path to the file to write to.
#      # Default: stdout
#      path: ./path/to/output.json
#    # Prints issues in columns representation separated by tabulations.
#    tab:
#      # Output path can be either stdout, stderr or path to the file to write to.
#      # Default: stdout
#      path: ./path/to/output.txt
#      # Print linter name in the end of issue text.
#      # Default: true
#      print-linter-name: true
#      # Use colors.
#      # Default: true
#      colors: false
#    # Prints issues in an HTML page.
#    # It uses the Cloudflare CDN (cdnjs) and React.
#    html:
#      # Output path can be either stdout, stderr or path to the file to write to.
#      # Default: stdout
#      path: ./path/to/output.html
#    # Prints issues in the Checkstyle format.
#    checkstyle:
#      # Output path can be either stdout, stderr or path to the file to write to.
#      # Default: stdout
#      path: ./path/to/output.xml
#    # Prints issues in the Code Climate format.
#    code-climate:
#      # Output path can be either stdout, stderr or path to the file to write to.
#      # Default: stdout
#      path: ./path/to/output.json
#    # Prints issues in the JUnit XML format.
#    junit-xml:
#      # Output path can be either stdout, stderr or path to the file to write to.
#      # Default: stdout
#      path: ./path/to/output.xml
#      # Support extra JUnit XML fields.
#      # Default: false
#      extended: true
#    # Prints issues in the TeamCity format.
#    teamcity:
#      # Output path can be either stdout, stderr or path to the file to write to.
#      # Default: stdout
#      path: ./path/to/output.txt
#    # Prints issues in the SARIF format.
#    sarif:
#      # Output path can be either stdout, stderr or path to the file to write to.
#      # Default: stdout
#      path: ./path/to/output.json

  # Add a prefix to the output file references.
  # Default: ""
  path-prefix: ""

  # Order to use when sorting results.
  # Possible values: file, linter, and severity.
  #
  # If the severity values are inside the following list, they are ordered in this order:
  #   1. error
  #   2. warning
  #   3. high
  #   4. medium
  #   5. low
  # Either they are sorted alphabetically.
  #
  # Default: ["linter", "file"]
  sort-order:
    - linter
    - severity
    - file # filepath, line, and column.

  # Show statistics per linter.
  # Default: true
  show-stats: false


# Options for analysis running.
run:
  # Timeout for total work, e.g. 30s, 5m, 5m30s.
  # If the value is lower or equal to 0, the timeout is disabled.
  # Default: 0 (disabled)
  timeout: 5m
  formatters:
    enable:
      - gci  # Moved from linters to formatters

  # The mode used to evaluate relative paths.
  # It's used by exclusions, Go plugins, and some linters.
  # The value can be:
  # - gomod: the paths will be relative to the directory of the go.mod file.
  # - gitroot: the paths will be relative to the git root (the parent directory of .git).
  # - cfg: the paths will be relative to the configuration file.
  # - wd (NOT recommended): the paths will be relative to the place where golangci-lint is run.
  # Default: cfg
  relative-path-mode: gomod

  # Exit code when at least one issue was found.
  # Default: 1
  issues-exit-code: 2

  # Include test files or not.
  # Default: true
  tests: false

  # List of build tags, all linters use it.
  # Default: []
  build-tags:
    - mytag

  # If set, we pass it to "go list -mod={option}". From "go help modules":
  # If invoked with -mod=readonly, the go command is disallowed from the implicit
  # automatic updating of go.mod described above. Instead, it fails when any changes
  # to go.mod are needed. This setting is most useful to check that go.mod does
  # not need updates, such as in a continuous integration and testing system.
  # If invoked with -mod=vendor, the go command assumes that the vendor
  # directory holds the correct copies of dependencies and ignores
  # the dependency descriptions in go.mod.
  #
  # Allowed values: readonly|vendor|mod
  # Default: ""
  modules-download-mode: readonly

  # Allow multiple parallel golangci-lint instances running.
  # If false, golangci-lint acquires file lock on start.
  # Default: false
  allow-parallel-runners: true

  # Allow multiple golangci-lint instances running, but serialize them around a lock.
  # If false, golangci-lint exits with an error if it fails to acquire file lock on start.
  # Default: false
  allow-serial-runners: true

  # Define the Go version limit.
  # Default: use Go version from the go.mod file, fallback on the env var GOVERSION, fallback on 1.22.
  go: '1.23'

  # Number of operating system threads (GOMAXPROCS) that can execute golangci-lint simultaneously.
  # Default: 0 (automatically set to match Linux container CPU quota and
  # fall back to the number of logical CPUs in the machine)
  concurrency: 4


severity:
  # Set the default severity for issues.
  #
  # If severity rules are defined and the issues do not match or no severity is provided to the rule
  # this will be the default severity applied.
  # Severities should match the supported severity names of the selected out format.
  # - Code climate: https://docs.codeclimate.com/docs/issues#issue-severity
  # - Checkstyle: https://checkstyle.sourceforge.io/property_types.html#SeverityLevel
  # - GitHub: https://help.github.com/en/actions/reference/workflow-commands-for-github-actions#setting-an-error-message
  # - TeamCity: https://www.jetbrains.com/help/teamcity/service-messages.html#Inspection+Instance
  #
  # @linter can be used as severity value to keep the severity from linters (e.g. revive, gosec, ...)
  #
  # Default: ""
  default: error

  # When a list of severity rules are provided, severity information will be added to lint issues.
  # Severity rules have the same filtering capability as exclude rules
  # except you are allowed to specify one matcher per severity rule.
  #
  # @linter can be used as severity value to keep the severity from linters (e.g. revive, gosec, ...)
  #
  # Only affects out formats that support setting severity information.
  #
  # Default: []
  rules:
    - linters:
        - dupl
      severity: info


  exclude-rules:
    - source: "(noinspection|TODO)"
      linters: [ godot ]
    - source: "//noinspection"
      linters: [ gocritic ]
    - path: "internal/controllers/grpc*"
      linters:
        - revive
    - path: "_test\\.go"
      linters:
        - bodyclose
        - dupl
        - funlen
        - goconst
        - gosec
        - noctx
        - wrapcheck
    - path: "cmd/client/client.go"
      linters:
        - mnd
        - govet
    - path: "internal/config/config.go"
      linters:
        - nosprintfhostport
        - funlen
`

# setup


# decisions

## router - mux: recommendation by the assessment
| Area      | Decision  |       Reason                  |
| ---------:| ---------:| ------:                       |
| Router    | ServeMux  | recommended by the assessment |

# test

- Unit Test -> for internal package (reusable components)
- Integration Test -> for external systems
- E2E Test -> for cmd package (entrypoint)
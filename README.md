# AI Coding Assistants Comparison

We believe AI coding tools will change the software creation process in huge ways. To understand which tools to use and where the industry stands, we conduct evaluations of these tools in a comparable way. If you just want to see the results, skip to the **Comparison** section.

## What We Evaluate

While AI coding tools often advertise their functionalities like "inline edits" or "large context," these are not direct coding features. We identified specific aspects of coding and tested each tool to see how they behave in these categories.

There are two main categories:
* **Integration with Human Workflow**
  * Evaluates the ease of integrating the AI tool into a typical human developer workflow.
* **AI Coding Score**
  * Evaluates aspects like handling single files, multiple files, documentation, and following architecture patterns.

For the detailed list, see this wiki page: [Evaluation Template 1.0](https://github.com/Venturing-Intellect/code-assist-comp/wiki/Evaluation-Template-1.0).

## Why Should You Trust Us?

* The evaluators are experienced software engineers who actively program daily and use various AI coding tools regularly.
* We use the same project for all tools to provide a consistent basis for comparison.
  * Project details: [Comparison Project](https://github.com/Venturing-Intellect/code-assist-comp/wiki/Software-project-for-comparison)

## Comparison

Last updated: 2024-12-23

### The Best Overall

**Cursor** ([https://www.cursor.com/](https://www.cursor.com/))  
Cursor is well-integrated into IDEs with convenient functionalities. It can use different language models, with Claude 3.5 as the default. It offers very powerful context capabilities. While there is room for improvement, Cursor outshines other AI coding assistants.

[Cursor Full Review](https://github.com/Venturing-Intellect/code-assist-comp/wiki/Evaluation-of-Cursor:-Java-project)

The only downside is that it's based on VSCode. IntelliJ or other IDE users face a steep learning curve. If that's a concern, read on.

### Runner-Up

**GitHub Copilot** ([https://github.com/features/copilot](https://github.com/features/copilot))  
Probably the most famous AI coding assistant. It integrates seamlessly into IDEs. This strength can also be a limitation, as there's no way to select which model is used. Generated code quality reflects this limitation, as better models are unavailable. However, it provides a productivity boost with significant room for improvement.

[Copilot Full Review](https://github.com/Venturing-Intellect/code-assist-comp/wiki/Evaluation-of-Copilot)

### Special Mention: Best Planning Tool

**GitHub Copilot Workspace** ([https://githubnext.com/projects/copilot-workspace](https://githubnext.com/projects/copilot-workspace))  

* It's not directly comparable to the previously evaluated solutions, as it's a web-based software project management tool. However, it generates code and works with your project, qualifying it as an AI coder.
* It's NOT a chatbot inside your IDE—you would still use your favorite chatbot in addition to Workspaces.
* It’s a tool for planning work before diving into the IDE to finish coding.

[GitHub Copilot Workspace Full Review](https://github.com/Venturing-Intellect/code-assist-comp/wiki/Evaluation-of-Github-Copilot-Workspace)

## Comparison Table

| Assistant                     | Overall Score | AI Coding Score | Integration with Human Workflow | Notes                                      |
|-------------------------------|---------------|------------------|----------------------------------|--------------------------------------------|
| Cursor (non-agentic)          | 8.45          | 6.9              | 10.0 (or 8.0)                    | [1]                                        |
| Copilot                       | 7.5           | 5.0              | 10.0                             |                                            |
| Continue.dev                  | 6.7           | 5.5              | 8.0                              |                                            |
| Tabnine                       | 6.5           | 6.0              | 7.0                              |                                            |
| Amazon Q                      | 6.3           | 5.6              | 7.0                              |                                            |
| GitHub Copilot Workspace      | 6.3           | 6.5              | 6.0                              | [2]                                        |
| Codeium                       | 3.9           | 4.9              | 3.0                              |                                            |
| Replit (Java)                 | 4.0           | 7.0              | 1.0                              |                                            |

**Notes:**
1. Cursor does not have an IntelliJ plugin. Scores may vary for IntelliJ IDE users.
2. Integration with the workflow for GitHub Copilot Workspace heavily depends on whether your issues are managed in GitHub and if you use VSCode. If not, this tool may not suit your needs.

To see all evaluation pages, visit our wiki: [Evaluation Wiki](https://github.com/Venturing-Intellect/code-assist-comp/wiki)
